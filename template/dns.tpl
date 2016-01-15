<!-- BEGIN: dns -->
<div class="container">
	<h2>DNS Records</h2>
	<p>You can add, delete, and manage your dns records below. DNS records are used to lookup information about a domain. Whenever you create a website, it will automatically add records.</p>
    <p>
    	<!-- BUTTON -->
    	
    	<span>
    		<a class="btn btn-sm btn-warning" href="#" onclick="javascript:$('#new_domain').removeClass('hidden')">
    			<span class="octicon octicon-book"></span>
    			<span class="normalize-text">Add New Domain</span>
    		</a>
    	</span>
    	<!-- /BUTTON -->
    </p>	
	
	<div id="new_domain" class="hidden">
		<form method="post" action="/dns/domain/add" id="newdomain">
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
        <div class="panel-heading">DNS Records</div>
        	<table class="table">
        			
        		<thead>
        			<tr>
        				<th>id</th>
        				<th class="col-md-3">Record Name</th>
        				<th class="col-md-1">Type</th>
        				<th class="col-md-3">Content</th>
        				<th class="col-md-1">Order</th>
        				<th class="col-md-1">Priority</th>
        				<th class="col-md-1">TTL</th>
        				<th class="col-md-2">Operations</th>
        			</tr>
        		</thead>
        		<tbody>
        			
        			<!-- BEGIN: domain -->
        			<tr>
        				<td colspan="8">
        				        Records for: <strong>{domain_name}</strong>

                    		<a class="btn btn-xs btn-success" href="#" onclick="javascript:$('#new_record_{record_domain_id}').removeClass('hidden')">
                    			<span class="octicon octicon-plus"></span> Add New Record
                    		</a>
                    		<a class="btn btn-xs btn-danger" href="/dns/domain/delete?domain={domain_name}" onclick="javascript:return confirm('Proceed with deleting {domain_name} and all related records?')">
                    			<span class="octicon octicon-trashcan"></span> Delete Domain
                    		</a>

        				</td>
        			</tr>
        			
        			
		            <form method="post" action="/dns/record/add" id="form_new_record_{record_domain_id}">
        			<input type="hidden" name="domain" value="{domain_name}">
        			<tr class="hidden" id="new_record_{record_domain_id}">
        				<td>
        				        New
        				</td>
        				<td>
        				        <input type="text" name="name" class="form-control" value="{domain_name}">
        				</td>
        				
        				<td>
        				        <select name="type" class="form-control">
        				            <option>A</option>
        				            <option>AAAA</option>
        				            <option>CNAME</option>
        				            <option>MX</option>
        				            <option>PTR</option>
        				            <option>SOA</option>
        				            <option>SRV</option>
        				            <option>NS</option>
        				        </select>
        				</td>
        				
        				<td>
        				        <input type="text" name="content" class="form-control" value="">
        				</td>
        				
        				<td>
        				        <input type="text" name="order" class="form-control" value="0">
        				</td>
        				
        				
        				<td>
        				        <input type="text" name="priority" class="form-control" value="0">
        				</td>
        				
        				<td>
        				        <input type="text" name="ttl" class="form-control" value="300">
        				</td>
        				
        				<td>
                    		<span>
                        		<button class="btn btn-xs btn-success" form="form_new_record_{record_domain_id}">
                        			<span class="octicon octicon-plus"></span>
                        		</button>
                        	</span>
                            
                    		<a class="btn btn-xs btn-danger" href="#" onclick="javascript:$('#new_record_{record_domain_id}').addClass('hidden')">
                    			<span class="octicon octicon-x"></span>
                    		</a>

        				</td>
        			</tr>
        			</form>
        			
        			
        			<!-- BEGIN: record -->
        			<form method="post" action="/dns/record/edit" id="form_edit_record_{record_id}">
        			<input type="hidden" name="record_id" value="{record_id}">
        			<tr>
        				<td>
        				        {record_id}
        				</td>
        				<td>
        				        <input type="text" name="name" class="form-control" value="{record_name}">
        				</td>
        				
        				<td>
        				        <select name="type" class="form-control">
        				            <option>{record_type}</option>
        				            <option>A</option>
        				            <option>AAAA</option>
        				            <option>CNAME</option>
        				            <option>MX</option>
        				            <option>PTR</option>
        				            <option>SOA</option>
        				            <option>SRV</option>
        				        </select>
        				</td>
        				
        				<td>
        				        <input type="text" name="content" class="form-control" value="{record_content}">
        				</td>
        				
        				<td>
        				        <input type="text" name="order" class="form-control" value="{record_ordername}">
        				</td>
        				
        				
        				<td>
        				        <input type="text" name="priority" class="form-control" value="{record_prio}">
        				</td>
        				
        				<td>
        				        <input type="text" name="ttl" class="form-control" value="{record_ttl}">
        				</td>
        				
        				<td>
                            <span>
                        		<button class="btn btn-xs btn-success" form="form_edit_record_{record_id}">
                        			<span class="octicon octicon-pencil"></span> Save
                        		</button>
                        	</span>
                            
                    		<a class="btn btn-xs btn-danger" href="/dns/record/delete?record_id={record_id}" onclick="javascript:return confirm('Proceed with deleting record?')">
                    			<span class="octicon octicon-trashcan"></span> Delete
                    		</a>

        				</td>
        			</tr>
        			</form>
        			<!-- END: record -->
        			<!-- END: domain -->

        			
        	</tbody>
        </table>

    </div>
</div>
<!-- END: dns -->

