<!-- BEGIN: databases -->
<div class="container">
	<h2>Databases</h2>
	<p>You can add, delete, and manage your MySQL databases or MySQL users below. Databases are the backends for storing website information.</p>

    <p>
    <!-- BUTTON -->
    <span class="pull-right">
    	<a class="btn btn-xs btn-info" href="{phpmyadmin_url}" target="_blank">
    		<span class="octicon octicon-browser"></span>
    		<span class="normalize-text">Go to phpMyAdmin</span>
    	</a>
    </span>
    <!-- /BUTTON -->
    </p>


    <div class="panel panel-default">
        <!-- Default panel contents -->
        <div class="panel-heading">Databases</div>
        	<table class="table">
        		<thead>
        			<tr>
        				<th class="col-md-3">Database Name</th>
        				<th>Operations / Remove Grants</th>
        			</tr>
        		</thead>
        		<tbody>
        			
        			
        			<!-- BEGIN: database -->
        			<tr>
        				<td>
        				        <p>{db_name}</p>
        				</td>
        				<td>

                            <div class="btn-group">
                                <div class="btn-group">
                                    <button type="button" class="btn btn-xs btn-default" data-toggle="dropdown">
                                    Grant Access <span class="caret"></span></button>
                                    <ul class="dropdown-menu" role="menu">
                                        <!-- BEGIN: add_grant -->
                                        <li><a href="/databases/grants/add?db_user={db_user}&db_name={db_name}">{db_user}</a></li>
                                        <!-- END: add_grant -->
                                    </ul>
                                </div>
                            </div>
                            
                    		<a class="btn btn-xs btn-danger" href="/databases/delete?db_name={db_name}">
                    			<span class="octicon octicon-trashcan"></span> Delete
                    		</a>
                    		
                    		<!-- BEGIN: grant -->
                    		    <a class="btn btn-xs btn-danger" href="/databases/grants/delete?db_user={db_user}&db_name={db_name}"><span class="octicon octicon-x"></span> {db_user}</a>
                    		<!-- END: grant -->
        				</td>
        			</tr>
        			<!-- END: database -->
        			
        			
        			
        			
        			<tr>
        				<td colspan="2">
                    		<form method="post" action="/databases/add" id="newdb">
                        		<div class="col-md-4"><input type="text" name="db_name" placeholder="Database Name" class="form-control"></div>
                        		
                            	<span>
                            		<button class="btn btn-sm btn-default" form="newdb">
                            			<span class="octicon octicon-database"></span>
                            			<span class="normalize-text">Add New Database</span>
                            		</button>
                            	</span>
                    		</form>
        				</td>
        			</tr>
        			
        			
        	</tbody>
        </table>

        <!-- Default panel contents -->
        <div class="panel-heading">Database Users</div>
        	<table class="table">
        		<thead>
        			<tr>
        				<th class="col-md-3">Username</th>
        				<th>Operations</th>
        			</tr>
        		</thead>
        		<tbody>
        		
        			<!-- BEGIN: user -->
        			<tr>
        				<td>{db_user}</td>
        				<td>
                            
                    		<div class="col-md-4">
                    		    <form method="post" action="/databases/users/edit" id="{db_user}_password">
                    		        <input type="hidden" name="db_user" value="{db_user}">
                    		        <input type="password" name="password" placeholder="New Password" class="form-control">
                    		    </form>
                    		</div>
                        	<span>
                        		<button class="btn btn-sm btn-info" form="{db_user}_password">
                        			<span class="octicon octicon"></span>
                        			<span class="normalize-text">Change Password</span>
                        		</button>
                        	</span>

                    		<a class="btn btn-sm btn-danger" href="/databases/users/delete?db_user={db_user}">
                    			<span class="octicon octicon-trashcan"></span>
                    		</a>
                    		
        				</td>
        			</tr>
        			<!-- END: user -->
        			
        	
        			<tr>
        				<td colspan="2">
                    		<form method="post" action="/databases/users/add" id="newuser">
                        		<div class="col-md-4"><input type="text" name="db_user" placeholder="New Username" class="form-control"></div>
                        		<div class="col-md-4"><input type="password" name="password" placeholder="New Password" class="form-control"></div>
                        		
                            	<span>
                            		<button class="btn btn-sm btn-default" form="newuser">
                            			<span class="octicon octicon-person"></span>
                            			<span class="normalize-text">Add New DB User</span>
                            		</button>
                            	</span>
                    		</form>
        				</td>
        			</tr>		
        			
        			
        			
        	</tbody>
        </table>
        
    </div>
    
</div>
<!-- END: databases -->

