<!-- BEGIN: websites -->
<div class="container">
	<h2>Websites</h2>
	<p>You can add, delete, and manage your websites below.</p>

	<!-- BUTTON -->
    <p>
    <form method="post" action="/websites/add" id="newsite">
    	<div class="col-md-4">
    	    <input type="text" name="domainname" class="form-control" placeholder="Domain name: example.com">
    	</div>
        <button type="submit" form="newsite" class="btn btn-sm btn-success">
    		<span class="octicon octicon-globe"></span>
    		<span class="normalize-text">Add New Site</span>
    	</button>
    </form>
    </p>
	<!-- /BUTTON -->




    <div class="panel panel-default">
        <!-- Default panel contents -->
        <div class="panel-heading">Configured Domains</div>
        	<table class="table">
        		<thead>
        			<tr>
        				<th class="col-md-2">Domain Name</th>
        				<th class="col-md-6">Document Root</th>
        				<th>Operations</th>
        			</tr>
        		</thead>
        		<tbody>
        			
        			<!-- BEGIN: domain -->
        			<tr>
        				<td>{domain}</td>
        				<td>{documentroot}</td>
        				<td>
                    		<a class="btn btn-sm btn-warning" href="/websites/sslmanager?vhost_id={vhost_id}">
                    			<span class="octicon octicon-shield"></span>
                    		</a>
                    		<a class="btn btn-sm btn-success" href="/filemanager?path={filemanager_path}">
                    			<span class="octicon octicon-file-directory"></span>
                    		</a>
                    		<a class="btn btn-sm btn-danger" href="/websites/delete?vhost_id={vhost_id}" onclick="javascript:return confirm('Delete {domain}?')">
                    			<span class="octicon octicon-trashcan"></span>
                    		</a>
        				</td>
        			</tr>
        			<!-- END: domain -->
        			
        	</tbody>
        </table>
    </div>


</div>
<!-- END: websites -->

