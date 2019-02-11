<!doctype html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <title>用户登陆页面</title>
    </head>
    <body>
        <form action="login.php" method="post">
            <p>email:<input type="text" name="email"></p>
            <p>password: <input type="text" name="password"></p>
            <p><input type="submit" name="submit" value="login"></p>
        </form>

        <?php 

           // 0. define mysql connect constant
           $servername = "localhost";
           $username = "newuser";
           $password = "testuser";
           $dbname = "assignment2";

           // 1. read data from input
           $email=$_POST["email"]??'';
           $password=$_POST["password"]??'';


           $conn = new mysqli($servername, $username, "", $dbname);
           // Check connection
           if ($conn->connect_error) {
               die("Connection failed: " . $conn->connect_error);
           } 
           
           $sql = "select * from User where email='".$email."' and password='".$password."';";
        //    echo $sql;
           $result = $conn->query($sql);

           if ($result->num_rows > 0) {
               echo "User exists; Login successfully";
           } else {
               echo "Empty or wrong input;\r\n";
               echo "Login Uncessfullly";
           }

$conn->close();
        ?>
    </body>
</html>