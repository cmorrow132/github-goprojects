<html>
<head>	
	<title>{{.Title}}</title>

	<script type="text/javascript" src="https://ajax.googleapis.com/ajax/libs/jquery/3.1.1/jquery.min.js"></script>
	<script type="text/javascript">
		$(document).ready(function() {
			$('#reloadData').click(function() {
				$('#dbResult').load('/' #main',function() {});
			});
		});
	</script>
</head>
<body>
<center>
	<h2><u>{{.Title}}</u></h2>
	<br>

	<div id="dbResult">
		<table width="80%" border=0>
		<tr>
		    <form method=POST action="/">
		        <input type=hidden name='cmd' value='filter'>
		        <td colspan=3><input type=text name='filter' size=50><input type=submit value='Filter'>&nbsp&nbsp&nbsp;<font size="-1"><a href="/">(Reset filter)</a></font></td>
		    </form>
		</tr><tr>
			<td colspan=3 bgcolor="#FFFC86" align=left border=0><font size="2"><b>{{.RecordCount}} matched records</b><br></td>
		</tr><tr>
			<th align=left><u><font size="3">First Name</th><th align=left><u><font size="3">Last Name</th><th align=left><u><font size="3">Phone</th>
		{{populateData}}
		<tr>
			<form method=POST action="/">
			    <input type=hidden name='cmd' value='add'>
				<td><input type=text name='firstname'></td><td><input type=text name='lastname'></td><td><input type=text name="phone"><input type=submit value='Add'></td>
			</form>
		</tr></table>
	</div>
</center>
<br>
<br><br>
<h4>Debug data</h4>
{{.DebugData}}
</body>
</html>

