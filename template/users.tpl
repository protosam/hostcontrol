<!-- BEGIN: users -->
<div class="container">
	<h2>Users</h2>
	<p>You can add, delete, and manage your system users below. These users will have access to create websites, email accounts, ftp users, and DNS records. They are not able to access your account.</p>


	<p>
		<form method="post" action="/users/add" id="newuser">
		    <div class="row">
        		<div class="col-md-4"><input type="text" name="username" placeholder="New Username" class="form-control"></div>
        		<div class="col-md-4"><input type="password" name="password" placeholder="New Password" class="form-control"></div>
            	<span>
            		<button class="btn btn-sm btn-default" form="newuser">
            			<span class="octicon octicon-person"></span>
            			<span class="normalize-text">Add New User</span>
            		</button>
            	</span>
        	</div>
        	<div>
        	    <!-- BEGIN: perms_all -->
        	    <input type="checkbox" name="allperms" id="allperms" value="Y"> <label for="allperms">All Permissions</label>
        	    <!-- END: perms_all -->
        	    <!-- BEGIN: perms_websites -->
        	    <input type="checkbox" name="websites" id="websites" value="Y"> <label for="websites">Websites</label>
        	    <!-- END: perms_websites -->
        	    <!-- BEGIN: perms_databases -->
        	    <input type="checkbox" name="databases" id="databases" value="Y"> <label for="databases">Databases</label>
        	    <!-- END: perms_databases -->
        	    <!-- BEGIN: perms_mail -->
        	    <input type="checkbox" name="mail" id="mail" value="Y"> <label for="mail">Mail</label>
        	    <!-- END: perms_mail -->
        	    <!-- BEGIN: perms_dns -->
        	    <input type="checkbox" name="dns" id="dns" value="Y"> <label for="dns">DNS Records</label>
        	    <!-- END: perms_dns -->
        	    <!-- BEGIN: perms_ftpusers -->
        	    <input type="checkbox" name="ftpusers" id="ftpusers" value="Y"> <label for="ftpusers">FTP Users</label>
        	    <!-- END: perms_ftpusers -->
        	    <!-- BEGIN: perms_sysusers -->
        	    <input type="checkbox" name="sysusers" id="sysusers" value="Y"> <label for="sysusers">Hostcontrol Users</label>
        	    <!-- END: perms_sysusers -->
        	</div>
		</form>
	</p>
	


    <div class="panel panel-default">
        <!-- Default panel contents -->
        <div class="panel-heading">Users</div>
        	<table class="table">
        		<thead>
        			<tr>
        				<th>Username</th>
        				<th>Email Address</th>
        				<th>Owner</th>
        				<th>Operations</th>
        			</tr>
        		</thead>
        		<tbody>
        			
        			<!-- BEGIN: user -->
        			<tr>
        				<td>{system_username}</td>
        				<td>{email_address}</td>
        				<td>{owned_by}</td>
        				<td>
                    		<a class="btn btn-sm btn-danger" href="/users/delete?username={system_username}" onclick="javascript:return confirm('Delete {system_username}?')">
                    			<span class="octicon octicon-trashcan"></span>
                    		</a>
                    		<a class="btn btn-sm btn-info" href="/users/sudo?username={system_username}" onclick="javascript:return confirm('Switch to {system_username}?')">
                    			<span class="octicon octicon-sign-in"></span>
                    		</a>
        				</td>
        			</tr>
        			<!-- END: user -->
        			
        	</tbody>
        </table>
    </div>

</div>
<!-- END: users -->