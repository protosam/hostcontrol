package routes

import (
	"fmt"
	"gopkg.in/macaron.v1"
	"github.com/protosam/vision"
	"github.com/protosam/hostcontrol/util"

	"os"
	"io/ioutil"


	"os/user"
	"path"
	"html"
	"strings"
	"strconv"
)


func init(){
	route("/filemanager", filemanager, "BOTH")
	route("/fileeditor", file_editor, "BOTH")
}


func filemanager(ctx *macaron.Context) (string){
	hcuser, auth := util.Auth(ctx, "any")
	if ! auth {
		ctx.Redirect("/", 302)
		return ""
	}


	var tpl vision.New
	tpl.TemplateFile("template/filemanager.tpl")


	suser, err := user.Lookup(hcuser.System_username)

	if err != nil {
		return die(ctx, string(err.Error()))
	}


	uid, err := strconv.Atoi(suser.Uid)
	if err != nil {
		return die(ctx, string(err.Error()))
	}


	gid, err := strconv.Atoi(suser.Gid)
	if err != nil {
		return die(ctx, string(err.Error()))
	}

	selected_object := path.Clean(util.Query(ctx, "path"))

	full_object := path.Clean(suser.HomeDir + "/" + selected_object)

    // check ownership...
    if ! util.ChkPerms(full_object, uid, gid) {
        return die(ctx, "You do not have access to object " + full_object)
    }
    

	delete_objectin := util.Query(ctx, "delete")
	delete_object := path.Clean(util.Query(ctx, "delete"))
	delete_object = path.Clean(suser.HomeDir + "/" + delete_object)

	if delete_objectin != "" {
		os.RemoveAll(delete_object)
	}

	newdirin := util.Query(ctx, "dirname")
	newdir := path.Clean(util.Query(ctx, "dirname"))
	newdir = path.Clean(full_object + "/" + newdir)

	newfilein := util.Query(ctx, "filename")
	newfile := path.Clean(util.Query(ctx, "filename"))
	newfile = path.Clean(full_object + "/" + newfile)


	if newdirin != "" {
		os.Mkdir(newdir, 0755)
		os.Chown(newdir, uid, gid)
	}

	if newfilein != "" {
		f, _ := os.Create(newfile)
		f.Close()
		os.Chown(newfile, uid, gid)
		os.Chmod(newfile, 0644)
	}
	

	tpl.GAssign("path_up", path.Dir(selected_object))
	tpl.GAssign("current_path", full_object)
	tpl.GAssign("selected_path", selected_object)

	objects, err := ioutil.ReadDir(full_object)

	if err != nil {
		return die(ctx, string(err.Error()))
	}

	tpl.Parse("filemanager")


	for _, file := range objects {
		tpl.Assign("filename", file.Name())

		mode := string(fmt.Sprintf("%s",file.Mode()))
		tpl.Assign("mode", mode)
		if file.IsDir() {
			tpl.Parse("filemanager/directory")
		}else{
			tpl.Parse("filemanager/file")
		}
	}

	return header(ctx) + tpl.Out() + footer(ctx)
}

func file_editor(ctx *macaron.Context) (string){
	hcuser, auth := util.Auth(ctx, "any")
	if ! auth {
		ctx.Redirect("/", 302)
		return ""
	}

	suser, err := user.Lookup(hcuser.System_username)

	if err != nil {
		return die(ctx, string(err.Error()))
	}


	selected_object := path.Clean(util.Query(ctx, "path"))
	full_object := path.Clean(suser.HomeDir + "/" + selected_object)


    // check ownership...
    uid, _ := strconv.Atoi(suser.Uid)
    gid, _ := strconv.Atoi(suser.Gid)
    if ! util.ChkPerms(full_object, uid, gid) {
        return die(ctx, "You do not have access to object " + full_object)
    }


	filecontents := util.Query(ctx, "filecontents")
	if filecontents != "" {
		filecontents = strings.Replace(filecontents, "\r\n", "\n", -1)
		ioutil.WriteFile(full_object, []byte(filecontents), 0644)
	}

	rawcontents, err := ioutil.ReadFile(full_object)
	if err != nil {
		return die(ctx, string(err.Error()))
	}


	content := html.EscapeString(string(rawcontents))
	


	var tpl vision.New
	tpl.TemplateFile("template/file-editor.tpl")

	tpl.Assign("path_up", path.Dir(selected_object))
	tpl.Assign("selected_path", selected_object)
	tpl.Assign("current_path", full_object)
	tpl.Assign("filedata", content)

	tpl.Parse("file-editor")


	return header(ctx) + tpl.Out() + footer(ctx)
}

