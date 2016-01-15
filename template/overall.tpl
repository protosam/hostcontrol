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
	<!-- begin navbar -->
	<nav class="navbar navbar-default navbar-fixed-top">
		<div class="container">
			<div class="navbar-header">
				<button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false" aria-controls="navbar">
					<span class="sr-only">Toggle navigation</span>
					<span class="icon-bar"></span>
					<span class="icon-bar"></span>
					<span class="icon-bar"></span>
				</button>
			<a class="navbar-brand" href="/dashboard"><span class="mega-octicon octicon-beaker"></span> Hostcontrol</a>
			</div>
			<!--.nav-collapse -->
			<div id="navbar" class="navbar-collapse collapse">
				<ul class="nav navbar-nav">
				<!-- removed the left nav -->
				</ul>
				<ul class="nav navbar-nav navbar-right">

					<li><a>Logged in as {username}</a></li>

					<li><a href="/logout"><span class="badge"><span class="octicon octicon-sign-out"></span> Logout</span></a></li>
				</ul>
			</div>
			<!--/.nav-collapse -->
		</div>
	</nav>
	<!-- end navbar -->

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
			<p><a href="http://just.ninja">Powered by Just.Ninja</a> &middot; Fireworks Release
		</footer>
	</div>

</body>
</html>
<!-- END: footer -->
