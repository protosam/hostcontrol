<!-- BEGIN: sslmanager -->
<div class="container">
	<h2>SSL Manager</h2>
	<p>You can enable/disable an SSL for {domain} below.</p>

    <form method="post" action="/websites/sslmanager/update?vhost_id={vhost_id}" id="sslmanage">
        <h4>SSL Certificate</h4>
        <div class="row">
            <div class="col-md-8">
                <textarea class="form-control" rows="20" name="crt_data">{ssl_certificate}</textarea>
            </div>
        </div>
        <h4>SSL Key</h4>
        <div class="row">
            <div class="col-md-8">
                <textarea class="form-control" rows="20" name="key_data">{ssl_key}</textarea>
            </div>
        </div>
        <h4>SSL CA Bundle</h4>
        <div class="row">
            <div class="col-md-8">
                <textarea class="form-control" rows="20" name="crtca_data">{ssl_ca_certificate}</textarea>
            </div>
        </div>
    
        <input type="checkbox" name="enablessl" id="enablessl" value="Y" {ssl_enabled}> <label for="enablessl">Enable SSL Certificate</label>
    </form>

    <button type="submit" form="sslmanage" class="btn btn-sm btn-success">
		<span class="octicon octicon-shield"></span>
		<span class="normalize-text">Update SSL Settings</span>
	</button>
	<a href="/websites" class="btn btn-sm btn-danger">
		<span class="octicon octicon-x"></span>
		<span class="normalize-text">Cancel</span>
	</a>
</div>
<!-- END: sslmanager -->

