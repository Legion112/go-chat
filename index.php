<?php
$pdo = new PDO('mysql:host=127.0.0.1;dbname=gochat', 'root', 'example', [
   PDO::ATTR_ERRMODE => PDO::ERRMODE_EXCEPTION
]);

$pdo->exec("INSERT INTO message (`date`, `username`, `text`)  VALUES(123213, 'Max', 'Hello there');");
