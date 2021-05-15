
<!DOCTYPE html>
<html>
<head>
<style>
body{
  font-family: Arial, Helvetica, sans-serif;
  border-collapse: collapse;
  width: 800px;
}

table {
  border-collapse: collapse;
}

td, th {
  border: 1px solid black;
  padding-left: 10px;
  padding: 2px;
}

img {
width:100%;
}


td{
  width: 500px;
}

th {
  width: 200px;
  padding-top: 4px;
  padding-left: 10px;
  padding-bottom: 4px;
  text-align: left;
  background-color: darkgray;  
  color: black;
}


h1 {
  font-size: 48px;  
}

img {
width:100%;
}


.title {
  width: 100%;
  padding-top: 400px;
  padding-left: 10%;  
}

.revision{
  vertical-align: bottom;
  width: 100%;
  padding-left: 10%; 
  padding-top: 400px;
}

</style>
</head>

<body>
<div class="title">
  <h1>THREAT AND RISK ANALYSIS</h1>
  <h1>{{.Title}}</h1>
  <h2>{{.Customer}}</h2>
</div>


<div class="revision">
  <table>
    <tr>
      <th>Version</th>
      <td>{{.Version}}</td>
    </tr>

    <tr>
      <th>Date</th>
      <td>{{.Date}}</td>
    </tr>

    <tr>
      <th>Authors</th>
      <td>{{.Author.Name}}</td>
    </tr>


  </table>
</div>

</body>
</html>
