<?php

$servername = "localhost";
$username = "h_wang";
$pass = "A00431268";
$dbname = "h_wang";

$link = mysqli_connect($servername, $username, $pass, $dbname);
mysqli_select_db($link, $dbname);
$sql = 'SELECT * FROM `User`';
$result = mysqli_query($link,$sql);
 # Execute the SELECT Query
  if ($result->num_rows <= 0){
    echo 'Retrieval of data from Database Failed - #'.mysql_errno().': '.mysql_error();
  }else{
// how to show a table in php?
// reference:https://stackoverflow.com/questions/4331353/retrieve-data-from-db-and-display-it-in-table-in-php-see-this-code-whats-wron
      ?>

<a href="signup.php">sign up </a>
<br/>
<a  href="login.php"> login </a>
<br/>
<a href="display.php"> display </a>
<br/>
<br/>

<table border="2">
  <thead>
    <tr>
      <th>user</th>
      <th>email</th>
      <th>password</th>
      <th>phone</th>
    </tr>
  </thead>
  <tbody>
    <?php
      if( $result->num_rows<=0 ){
        echo '<tr><td colspan="4">No Rows Returned</td></tr>';
      }else{
        while( $row =$result->fetch_assoc() ){
          echo "<tr><td>{$row['user']}</td><td>{$row['email']}</td><td>***</td><td>{$row['phone']}</td></tr>\n";
        }
      }
    ?>
  </tbody>
</table>

<?php
  }

?>
