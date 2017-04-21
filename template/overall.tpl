<!-- BEGIN: header -->
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="utf-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<title>Hostcontrol</title>

	<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.min.css">
	<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap-theme.min.css">
	<link href="/resources/style.css" rel="stylesheet">
	<link href="/resources/octicons/octicons.css" rel="stylesheet">

	<script src="https://ajax.googleapis.com/ajax/libs/jquery/2.1.4/jquery.min.js"></script>
	<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/js/bootstrap.min.js"></script>
	<script src="/resources/ace/ace.js"></script>
	<script src="/resources/ace/ext-modelist.js"></script>

	<!-- HTML5 shim and Respond.js for IE8 support of HTML5 elements and media queries -->
	<!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
	<!--[if lt IE 9]>
		<script src="https://oss.maxcdn.com/html5shiv/3.7.2/html5shiv.min.js"></script>
		<script src="https://oss.maxcdn.com/respond/1.4.2/respond.min.js"></script>
	<![endif]-->
</head>
<body>
<div class="page-header row nopad">
  <div class="pull-right normalize-text pad5">Logged in as {username} <a href="/logout"><span class="badge"><span class="octicon octicon-sign-out"></span> Logout</span></a></div>
  <h1 class="pull-left nopad"><a href="/dashboard"><span class="mega-octicon octicon-beaker"></span> Hostcontrol</a></h1>
</div>

<ul class="nav nav-pills">
	<!-- BEGIN: websitesbtn -->
	<li role="presentation">
	    <a href="/websites">
			<span class="octicon octicon-globe"></span>
			<span class="normalize-text">Websites</span>
	    </a>
	</li>
	<!-- END: websitesbtn -->
	
	<!-- BEGIN: databasesbtn -->
	<li role="presentation">
	    <a href="/databases">
			<span class="octicon octicon-database"></span>
			<span class="normalize-text">Databases</span>
		</a>
	</li>
	<!-- END: databasesbtn -->
	
	<!-- BEGIN: dnsbtn -->
	<li role="presentation">
	    <a href="/dns">
			<span class="octicon octicon-repo"></span>
			<span class="normalize-text">DNS Records</span>
		</a>
	</li>
	<!-- END: dnsbtn -->
	
	<!-- BEGIN: mailbtn -->
	<li role="presentation">
	    <a href="/mail">
			<span class="octicon octicon-mail"></span>
			<span class="normalize-text">Email</span>
		</a>
	</li>
	<!-- END: mailbtn -->
	
	<!-- BEGIN: ftpusersbtn -->
	<li role="presentation">
	    <a href="/ftpusers">
			<span class="octicon octicon-file-symlink-directory"></span>
			<span class="normalize-text">FTP Users</span>
    	</a>
	</li>
	<!-- END: ftpusersbtn -->
	
		
	<!-- BUTTON -->
	<li role="presentation">
		<a href="/filemanager">
			<span class="octicon octicon-file-directory"></span>
			<span class="normalize-text">File Manager</span>
		</a>
	</li>
	<!-- /BUTTON -->


	<!-- BEGIN: firewallbtn -->
	<li role="presentation">
		<a href="/firewall">
			<span class="octicon octicon-flame"></span>
			<span class="normalize-text">Firewall</span>
		</a>
	</li>
	<!-- END: firewallbtn -->

	<!-- BUTTON -->
	<li role="presentation">
		<a href="{console_url}" target="_blank">
			<span class="octicon octicon-terminal"></span>
			<span class="normalize-text">Console</span>
		</a>
	</li>
	<!-- /BUTTON -->
	
	<!-- BEGIN: servicesbtn -->
	<li role="presentation">
		<a href="/services">
			<span class="octicon octicon-pulse"></span>
			<span class="normalize-text">Services</span>
		</a>
	</li>
	<!-- END: servicesbtn -->
	
	<!-- BEGIN: usersbtn -->
	<li role="presentation">
		<a href="/users">
			<span class="octicon octicon-organization"></span>
			<span class="normalize-text">Users</span>
		</a>
	</li>
	<!-- END: usersbtn -->

	<!-- BUTTON -->
	<li role="presentation">
		<a href="/settings">
			<span class="octicon octicon-tools"></span>
			<span class="normalize-text">Settings</span>
		</a>
	</li>
	<!-- /BUTTON -->
</ul>

					
					


	
	
	
	<!-- BEGIN: info -->
	<p class="alert alert-success">{message}</p>
	<!-- END: info -->
	<!-- BEGIN: error -->
	<p class="alert alert-warning">{message}</p>
	<!-- END: error -->
<!-- END: header -->
<!-- BEGIN: footer -->
    <div class="container marketing">
		<hr class="featurette-divider">
		<footer>
			<p class="pull-right"><a href="#">Back to top</a></p>
			<p>Fireworks Release
		</footer>
	</div>

</body>
</html>
<!-- END: footer -->
