<!doctype html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <title>User Signup Page</title>
    </head>
    <body>
        <form action="signup.php" method="post">
            <p>username:<input type="text" name="name"></p>
            <p>password: <input type="text" name="password"></p>
            <p>email: <input type="text" name="email"></p>
            <p>phone: <input type="text" name="phone"></p>
            <p><input type="submit" name="submit" value="register"></p>
        </form>

        <?php 

           // 0. define mysql connect constant
           $servername = "localhost";
           $username = "newuser";
           $password = "testuser";
           $dbname = "h_wang";

           // 1. read data from input
           $user=$_POST["name"]??'';
           $password=$_POST["password"]??'';
           $email=$_POST["email"]??'';
           $phone=$_POST["phone"]??'';
           
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
           $conn = new mysqli($servername, $username, "", $dbname);
           // Check connection
           if ($conn->connect_error) {
               die("Connection failed: " . $conn->connect_error);
           } 
           
           $sql = "INSERT INTO User (user, password,email,phone)
           VALUES ('$user', '$password', '$email','$phone')";
           
           if ($conn->query($sql) === TRUE) {
               echo "New record created successfully";
           } else {
               echo "Error: " . $sql . "<br>" . $conn->error;
           }

$conn->close();
        ?>
    </body>
</html>