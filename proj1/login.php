<?php
$secret_pass = "really_secure_password";

@session_start();

if (isset($_GET['pass']) && strcmp($_GET['pass'], $secret_pass) == 0) {
    $_SESSION['user_logged_in'] = 1;
} else {
    echo "Wrong pass!";
}
