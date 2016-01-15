<!-- BEGIN: filemanager -->
<div class="container">
	<h4>Broswing files </h4>


    <div class="panel panel-default">
        <!-- Default panel contents -->
        <div class="panel-heading">{current_path}</div>
        <form method="post" action="/filemanager?path={selected_path}" id="file">
        	<table class="table">
        		<thead>
        			<tr>
        				<th class="col-md-8">Filename</th>
        				<th>Permissions</th>
        				<th>Operations</th>
        			</tr>
        		</thead>
        		<tbody>
        			<tr>
        				<td><a href="/filemanager?path={path_up}"><span class="glyphicon glyphicon-level-up"></span> Up Directory</a></td>
        				<td></td>
        				<td></td>
        			</tr>
        			
        			
        			<!-- BEGIN: directory -->
        			<tr>
        				<td><span class="octicon octicon-file-directory"></span> <a href="/filemanager?path={selected_path}/{filename}">{filename}</a></td>
        				<td>{mode}</td>
        				<td>
                    		<a class="btn btn-sm btn-danger" href="/filemanager?path={selected_path}&delete={selected_path}/{filename}" onclick="javascript:return confirm('Delete {filename}?')">
                    			<span class="octicon octicon-trashcan"></span>
                    		</a>
        				</td>
        			</tr>
        			<!-- END: directory -->
        			
        			
        			<tr>
        				<td colspan="2"><input type="text" name="dirname" class="form-control" placeholder="New directory name"></td>
        				<td>
        				    <button type="submit" form="file" class="btn btn-sm btn-info">
                    			<span class="octicon octicon-file-directory"></span>
                    			<span class="normalize-text">Create Directory</span>
                    		</button>
        				</td>
        			</tr>
        			
        			
        			
        			<!-- BEGIN: file -->
        			<tr>
        				<td><span class="octicon octicon-file-text"></span> <a href="/fileeditor?path={selected_path}/{filename}">{filename}</a></td>
        				<td>{mode}</td>
        				<td>
                    		<a class="btn btn-sm btn-danger" href="/filemanager?path={selected_path}&delete={selected_path}/{filename}" onclick="javascript:return confirm('Delete {filename}?')">
                    			<span class="octicon octicon-trashcan"></span>
                    		</a>
        				</td>
        			</tr>
        			<!-- END: file -->
        			
        			

        			<tr>
        				<td colspan="2"><input type="text" name="filename" class="form-control" placeholder="New file name"></td>
        				<td>
                            <button type="submit" form="file" class="btn btn-sm btn-success">
                                <span class="octicon octicon-file-text"></span>
                                <span class="normalize-text">Create File</span>
                            </button>
        				</td>
        			</tr>
        			
        			
        		</tbody>
        	</table>
        </form>
    </div>

 
</div>

<!-- END: filemanager -->

