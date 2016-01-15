<!-- BEGIN: mail -->
<div class="container">
	<h2>Mail</h2>
	<p>You can add, delete, and manage your email accounts below. If you just created a website, you will have to add the domain here as well if you want to use email on this server.</p>
    <p>
    	<!-- BUTTON -->
    	<span>
    		<a class="btn btn-sm btn-info" href="#" onclick="javascript:$('#new_domain').removeClass('hidden')">
    			<span class="octicon octicon-globe"></span>
    			<span class="normalize-text">Add New Domain</span>
    		</a>
    	</span>
    	<!-- /BUTTON -->
    </p>

    <p>
    <!-- BUTTON -->
    <span class="pull-right">
    	<a class="btn btn-xs btn-info" href="{webmail_url}" target="_blank">
    		<span class="octicon octicon-browser"></span>
    		<span class="normalize-text">Go to Webmail</span>
    	</a>
    </span>
    <!-- /BUTTON -->
    </p>

	<div id="new_domain" class="hidden">
		<form method="post" action="/mail/domain/add" id="newdomain">
            <p>
        		<div class="col-md-4"><input type="text" name="domain" placeholder="example.com" class="form-control"></div>
            	<span>
            		<button class="btn btn-sm btn-success" form="newdomain">
            			<span class="octicon octicon-plus"></span>
            			<span class="normalize-text">Add New Domain</span>
            		</button>
            	</span>
        	</p>
		</form>
	</div>
	
    <div class="panel panel-default">
        <!-- Default panel contents -->
        <div class="panel-heading">Email Accounts</div>
        	<table class="table">
        			
        		<thead>
        			<tr>
        				<th class="col-md-3">Account</th>
        				<th>Operations</th>
        			</tr>
        		</thead>
        		<tbody>
        			
        			<!-- BEGIN: domain -->
        			<tr>
        				<td colspan="2">
        				        Accounts for: <strong>{domain_name}</strong>

                    		<a class="btn btn-xs btn-success" href="#" onclick="javascript:$('#add_new_email').removeClass('hidden')">
                    			<span class="octicon octicon-plus"></span> Add New Account
                    		</a>
                    		<a class="btn btn-xs btn-danger" href="/mail/domain/delete?domain={domain_name}" onclick="javascript:return confirm('Proceed with deleting {domain_name} and all related records?')">
                    			<span class="octicon octicon-trashcan"></span> Delete Domain
                    		</a>

        				</td>
        			</tr>
        			
        			<input type="hidden" name="domain" value="{domain_name}">
        			<tr class="hidden" id="add_new_email">
        				<td colspan="2">
                		    <form method="post" action="/mail/users/add" id="form_new_email">
                		        <div class="col-md-3"><input type="text" name="username" placeholder="Username" class="form-control"></div>
                		        <div class="col-md-3">
                        			<span class="normalize-text">@{domain_name} <input type="hidden" name="domain" value="{domain_name}"></span>
                		        </div>
                		        <div class="col-md-3"><input type="password" name="password" placeholder="New Password" class="form-control"></div>
                                <div class="col-md-3">
                                	<span>
                                		<button class="btn btn-xs btn-success" form="form_new_email">
                                			<span class="octicon octicon-plus"></span>
                                		</button>
                                	</span>
        
                            		<a class="btn btn-xs btn-danger" href="#" onclick="javascript:$('#form_new_email').addClass('hidden')">
                            			<span class="octicon octicon-x"></span>
                            		</a>
                                </div>
                		    </form>

        				</td>
        			</tr>
        			
        			
        			<!-- BEGIN: email -->
        			<tr>
        				<td>
        				        {email}
        				</td>
        				
        				<td>
                            
                    		<div class="col-md-4">
                    		    <form method="post" action="/mail/users/edit" id="form_edit_record_{email_id}">
                    		        <input type="hidden" name="email" value="{email}">
                    		        <input type="password" name="password" placeholder="New Password" class="form-control">
                    		    </form>
                    		</div>
                        	<span>
                        		<button class="btn btn-sm btn-info" form="form_edit_record_{email_id}">
                        			<span class="octicon octicon-pencil"></span>
                        			<span class="normalize-text">Change Password</span>
                        		</button>
                        	</span>

                    		<a class="btn btn-sm btn-danger" href="/mail/users/delete?email={email}">
                    			<span class="octicon octicon-trashcan"></span>
                    		</a>
                    		
        				</td>
        			</tr>
        			<!-- END: email -->
        			<!-- END: domain -->

        			
        	</tbody>
        </table>

    </div>
</div>
<!-- END: mail -->

