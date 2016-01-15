<!-- BEGIN: login -->
<!DOCTYPE html>
<html lang="en">
  <head>
        <meta charset="utf-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <title>Login</title>

        <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.min.css">
        <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap-theme.min.css">
        <link href="/resources/style.css" rel="stylesheet">
        <link href="/resources/octicons/octicons.css" rel="stylesheet">
        <link href="/resources/signin.css" rel="stylesheet">

        <script src="https://ajax.googleapis.com/ajax/libs/jquery/2.1.4/jquery.min.js"></script>
        <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/js/bootstrap.min.js"></script>

        <!-- HTML5 shim and Respond.js for IE8 support of HTML5 elements and media queries -->
        <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
        <!--[if lt IE 9]>
                <script src="https://oss.maxcdn.com/html5shiv/3.7.2/html5shiv.min.js"></script>
                <script src="https://oss.maxcdn.com/respond/1.4.2/respond.min.js"></script>
        <![endif]-->
  </head>

  <body>

    <div class="container">

      <form class="form-signin" method="post" action="/login">
        <h2 class="form-signin-heading">Please login below</h2>

	<!-- BEGIN: fail -->
	<div class="alert alert-danger" role="alert">
		<strong>Login Failed!</strong> You typed in a bad username/password combination.
	</div>
	<!-- END: fail -->
	<!-- BEGIN: logged_out -->
	<div class="alert alert-info" role="alert">
		<strong>Logged out!</strong> You have been logged out. If you wish, you can login again below.
	</div>
	<!-- END: logged_out -->

        <label for="inputEmail" class="sr-only">Username</label>
        <input type="text" id="inputEmail" class="form-control" name="username" placeholder="Username" required autofocus>
        <label for="inputPassword" class="sr-only">Password</label>
        <input type="password" id="inputPassword" class="form-control" name="password" placeholder="Password" required>
        <div class="checkbox">
          <label>
            <input type="checkbox" name="rememberme" value="checked"> Remember me
          </label>
        </div>
        <button class="btn btn-lg btn-primary btn-block" type="submit">Sign in</button>
      </form>

    </div> <!-- /container -->


    <!-- IE10 viewport hack for Surface/desktop Windows 8 bug -->
    <script src="/resources/js/ie10-viewport-bug-workaround.js"></script>
  </body>
</html>
<!-- END: login -->
