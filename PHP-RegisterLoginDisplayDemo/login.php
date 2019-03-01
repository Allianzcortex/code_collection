<!doctype html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <title>User Login Page</title>
    </head>

    <body>

<a href="signup.php">sign up </a>
<br/>
<a  href="login.php"> login </a>
<br/>
<a href="display.php"> display </a>
<br/>
<br/>



        <form action="login.php" method="post">
            <p>email:<input type="text" name="email"></p>
            <p>password: <input type="text" name="password"></p>
            <p><input type="submit" name="submit" value="login"></p>
        </form>

        <?php 

           // 0. define mysql connect constant
           $servername = "localhost";
           $username = "root";
           $pass = "root";
           $dbname = "root";

           // 1. read data from input
           $email=$_POST["email"]??'';
           $password=$_POST["password"]??'';


           $link = mysqli_connect($servername, $username, $pass, $dbname);

           //Check the connection. Give error message if any error 
           if (!$link) die("Couldn't connect to MySQL");
                      mysqli_select_db($link, $dbname);
           
           $sql = "select * from User where email='".$email."' and password='".$password."';";
        //    echo $sql;
        $result = mysqli_query($link,$sql);
           
        if (mysqli_num_rows($result) >0 ) {
            echo "User exists; Login successfully";
        } else {
            echo "Empty or wrong input;\r\n";
               echo "Login Uncessfullly";
        }

$conn->close();
        ?>
    </body>
</html>