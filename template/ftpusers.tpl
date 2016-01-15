<!-- BEGIN: ftpusers -->
<div class="container">
	<h2>FTP Users</h2>
	<p>You can add, delete, and manage your FTP users below. FTP users are meant to access specific domains on your account. Please note that even though the user is isolated to your specified directory, the code uploaded by the user is able to access any of the data for your other websites.</p>
    <p>
		<form method="post" action="/ftpusers/add" id="newuser">
    		<div class="col-md-2"><input type="text" name="ftpuser" placeholder="New Username" class="form-control"></div>
    		<div class="col-md-2"><input type="password" name="password" placeholder="New Password" class="form-control"></div>
		    <div class="col-md-4"><input type="text" name="homedir" value="{homedir}/" class="form-control"></div>

        	<span>
        		<button class="btn btn-sm btn-danger" form="newuser">
        			<span class="octicon octicon-person"></span>
        			<span class="normalize-text">Add New User</span>
        		</button>
        	</span>
		</form>
    </p>



    <div class="panel panel-default">
        <!-- Default panel contents -->
        <div class="panel-heading">FTP Users</div>
        	<table class="table">
        		<thead>
        			<tr>
        				<th>Username</th>
        				<th class="col-md-4">Home Directory</th>
        				<th>Operations</th>
        			</tr>
        		</thead>
        		<tbody>
        			<!-- BEGIN: user -->
        			<tr>
        				<td>{username}</td>
        				<td>{homedir}</td>
        				<td>
        				

                    		<div class="col-md-6">
                    		    <form method="post" action="/ftpusers/edit" id="{username}_password">
                    		        <input type="hidden" name="ftpuser" value="{username}">
                    		        <input type="password" name="password" placeholder="New Password" class="form-control">
                    		    </form>
                    		</div>
                        	<span>
                        		<button class="btn btn-sm btn-info" form="{db_user}_password">
                        			<span class="octicon octicon"></span>
                        			<span class="normalize-text">Change Password</span>
                        		</button>
                        	</span>

                    		<a class="btn btn-sm btn-danger" href="/ftpusers/delete?ftpuser={username}" onclick="javascript:return confirm('Delete {username}?')">
                    			<span class="octicon octicon-trashcan"></span>
                    		</a>
        				</td>
        			</tr>
        			<!-- END: user -->
        	</tbody>
        </table>
    </div>
    
    
</div>
<!-- END: ftpusers -->

