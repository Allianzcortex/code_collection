<?php

$servername = "localhost";
$username = "newuser";
$password = "testuser";
$dbname = "assignment2";

$conn = new mysqli($servername, $username, "", $dbname);
$selectSQL = 'SELECT * FROM `User`';
$result = $conn->query($selectSQL);
 # Execute the SELECT Query
  if ($result->num_rows <= 0){
    echo 'Retrieval of data from Database Failed - #'.mysql_errno().': '.mysql_error();
  }else{
      ?>
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
