<!-- BEGIN: file-editor -->
<div class="container">
	<h4>Editing file {current_path}</h4>
</div>

<div class="container">
	<div id="editor">{filedata}</div>
	<form method="post" action="/fileeditor?path={selected_path}" id="file">
		<input type="hidden" name="filecontents" value="">
	</form>
</div>


<div class="container">
	<!-- BUTTON -->
	<span id="changes_made" class="hidden">
		<button type="submit" form="file" class="btn btn-sm btn-success">
			<span class="octicon octicon-file-text"></span>
			<span class="normalize-text">Save File</span>
		</button>
	</span>
	<!-- /BUTTON -->

	<!-- BUTTON -->
	<span>
		<a class="btn btn-sm btn-danger" href="/filemanager?path={path_up}">
			<span class="octicon octicon-x"></span>
			<span class="normalize-text">Close</span>
		</a>
	</span>
	<!-- /BUTTON -->

</div>

<script>
	var filePath = "something.go"
	var editor = ace.edit("editor");
	var modelist = ace.require("ace/ext/modelist")
	var mode = modelist.getModeForPath(filePath).mode

	editor.setTheme("ace/theme/monokai");
	editor.getSession().setMode(mode);

	var input = $('input[name="filecontents"]');
	editor.getSession().on("change", function () {
		input.val(editor.getSession().getValue());
		$("#changes_made").removeClass("hidden")
	});
</script>
<!-- END: file-editor -->
