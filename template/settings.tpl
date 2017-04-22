<!-- BEGIN: settings -->
<div class="container">
	<h2>Settings</h2>
	<p>Manage the account settings for {username} below.</p>
</div>



<form method="post" action="/settings/update" id="update">
    <div class="container">
        <h4>Change Password</h4>
        <div class="row">
        	<div class="col-md-2"><input type="password" name="password" placeholder="Current Password" class="form-control"></div>
        	<div class="col-md-2"><input type="password" name="new_password" placeholder="New Password" class="form-control"></div>
        	<div class="col-md-2"><input type="password" name="new_password_verify" placeholder="Verify New Password" class="form-control"></div>
    	</div>
    	
    	<br>
    	<span>
    		<button class="btn btn-sm btn-default" form="update">
    			<span class="octicon octicon-gear"></span>
    			<span class="normalize-text">Update Settings</span>
    		</button>
    	</span>
    
        <hr class="featurette-divider">
    </div>
</form>


<div class="container">

	<h4>Tokens</h4>
	<p>Manage API tokens below.</p>



	<p>
		<form method="post" action="/settings/tokens/add" id="newtoken">
    		<div class="col-md-4"><input type="text" name="description" placeholder="New token description" class="form-control"></div>
    		
        	<span>
        		<button class="btn btn-sm btn-default" form="newtoken">
        			<span class="octicon octicon-key"></span>
        			<span class="normalize-text">Add New Token</span>
        		</button>
        	</span>
		</form>
	</p>

    <div class="panel panel-default">
        <!-- Default panel contents -->
        <div class="panel-heading">Users</div>
        	<table class="table">
        		<thead>
        			<tr>
        				<th>Token</th>
        				<th>Description</th>
        				<th>Operations</th>
        			</tr>
        		</thead>
        		<tbody>
        			
        			<!-- BEGIN: token -->
        			<tr>
        				<td>{token}</td>
        				<td>{description}</td>
        				<td>
                    		<a class="btn btn-sm btn-danger" href="/settings/tokens/delete?token={raw_token}" onclick="javascript:return confirm('Confirm delete for:\n {description}')">
                    			<span class="octicon octicon-trashcan"></span>
                    		</a>
        				</td>
        			</tr>
        			<!-- END: token -->
        			
        	</tbody>
        </table>
    </div>
</div>


<div class="container">

	<h3>API Calls</h3>
	<p>Below are the availabel API calls.</p>

    <h4></h4>
    <p>To use your token against the API, you must add the parameter <code>api_token=##/###TOKEN###STRING###</code> via the request query string or post data. Below are two examples with curl.</p>
    <p><code>curl --data "api_token=api_token=##/###TOKEN###STRING###" https://SERVER_IP_HERE/api/users/list</code></p>
    <p><code>curl https://SERVER_IP_HERE/api/users/list?api_token=api_token=##/###TOKEN###STRING###</code></p>
    
    <h4>Systems Info</h4>
    <p>These API calls will return information about the server ecosystem.</p>
    
    <p><code>/api/distro</code></p>
    <p><code>/api/ftp</code></p>
    <p><code>/api/dns</code></p>
    <p><code>/api/mail</code></p>
    <p><code>/api/sql</code></p>
    <p><code>/api/web</code></p>
    
    <hr class="featurette-divider">
    <h4>User Management</h4>
    
    <p><code>/api/users/list</code></p>
    <p><code>/api/users/add?username=&password=[&allperms=Y][&mail=Y][&sysusers=Y][&databases=Y][&ftpusers=Y][&websites=Y][&dns=Y]</code></p>
    <p><code>/api/users/delete?username=</code></p>
    
    <hr class="featurette-divider">
    <h4>FTP User Management</h4>
    
    <p><code>/api/ftpusers/list</code></p>
    <p><code>/api/ftpusers/add?ftpuser=&password=&homedir=</code></p>
    <p><code>/api/ftpusers/edit?username=?&password=</code></p>
    <p><code>/api/ftpusers/delete?ftpuser=</code></p>
    
    <hr class="featurette-divider">
    <h4>DNS Management</h4>
    
    <p><code>/api/dns/list</code></p>
    <p><code>/api/dns/domain/add?domain=</code></p>
    <p><code>/api/dns/domain/delete?domain=</code></p>
    <p><code>/api/dns/record/add?domain=&name=&content=&type=&ttl=&priority</code></p>
    <p><code>/api/dns/record/edit?record_id=&name=&content=&type=&ttl=&priority</code></p>
    <p><code>/api/dns/record/delete?record_id=</code></p>
    
    <hr class="featurette-divider">
    <h4>Mail Management</h4>
    
    <p><code>/api/mail/list</code></p>
    <p><code>/api/mail/domain/add?domain=</code></p>
    <p><code>/api/mail/domain/delete?domain=</code></p>
    <p><code>/api/mail/users/add?domain=&username=&password=</code></p>
    <p><code>/api/mail/users/edit?email=&password=</code></p>
    <p><code>/api/mail/users/delete?email=</code></p>
    
    <hr class="featurette-divider">
    <h4>SQL Management</h4>
    
    <p><code>/api/sql/databases/list</code></p>
    <p><code>/api/sql/databases/add?db_name=</code></p>
    <p><code>/api/sql/databases/delete?db_name=</code></p>
    
    <p><code>/api/sql/users/list</code></p>
    <p><code>/api/sql/users/add?db_user=&password</code></p>
    <p><code>/api/sql/users/edit?db_user=&password</code></p>
    <p><code>/api/sql/users/delete?db_user=</code></p>
    
    <p><code>/api/sql/grants/list</code></p>
    <p><code>/api/sql/grants/add?db_name=&db_user=</code></p>
    <p><code>/api/sql/grants/delete?db_name=&db_user=</code></p>
    
    <hr class="featurette-divider">
    <h4>Website Management</h4>
    
    <p><code>/api/web/domain/list</code></p>
    <p><code>/api/web/domain/add?domainname=</code></p>
    <p><code>/api/web/domain/delete?vhost_id=</code></p>
    <p><code>/api/web/domain/sslmanage?vhost_id=&enablessl=[Y|N]&crt_data=&crtca_data=&key_data</code></p>
</div>
<!-- END: settings -->

