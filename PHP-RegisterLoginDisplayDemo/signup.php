<!doctype html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <title>User Signup Page</title>
    </head>
    <body>

<a href="signup.php">sign up </a>
<br/>
<a  href="login.php"> login </a>
<br/>
<a href="display.php"> display </a>
<br/>
<br/>


        <form action="signup.php" method="post">
            <p>username: <input type="text" name="name"></p>
            <p>password: <input type="text" name="password"></p>
            <p>email: <input type="text" name="email"></p>
            <p>phone: <input type="text" name="phone"></p>
            <p><input type="submit" name="submit" value="register" onClick="location.href=location.href"></p>
        </form>

        <?php 

           // 0. define mysql connect constant
           $servername = "localhost";
           $username = "root";
           $pass = "root";
           $dbname = "root";

           // 1. read data from input
           $user=$_POST["name"]??'';
           $password=$_POST["password"]??'';
           $email=$_POST["email"]??'';
           $phone=$_POST["phone"]??'';

           
	   if($user==null){
	throw new Exception("username is null");
}
        //    CREATE TABLE IF NOT EXISTS User (
        //     id INT AUTO_INCREMENT,
        //     user varchar(50),
        //     password varchar(50),
        //     email varchar(50),
        //     phone varchar(50),
        //     PRIMARY KEY (id)
        // )  ENGINE=INNODB;

           // 2. write to database

           // Create connection
           // I don't know why ......
          $link = mysqli_connect($servername, $username, $pass, $dbname);

//Check the connection. Give error message if any error 
if (!$link) die("Couldn't connect to MySQL");
//It will print error message if there is any error.
//mysqli_select_db() is used to select the database 
mysqli_select_db($link, $dbname)
	
	or die("Couldn't open $db: ".mysqli_error($link));
            $sql = "INSERT INTO User (user, password,email,phone)
           VALUES ('$user', '$password', '$email','$phone')";
$result = mysqli_query($link,$sql);
           
           if ($result === TRUE) {
               echo "User Register  successfully";
           } else {
               echo "Error: " . $sql . "<br>" . $conn->error;
           }

$conn->close();
        ?>
    </body>
</html>