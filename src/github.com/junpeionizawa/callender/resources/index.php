<!DOCTYPE html>
<html lang="en">

<head>
	<meta charset="UTF-8">
	<title>テスト</title>
	<link rel="stylesheet" type="text/css" href="../resources/styles.css">
</head>
<body>
	<div class="calenderbox">
		<pre><h1 class="month_now"><span>{{.Year }}</span></h1></pre>
		<div id ="box1"> 
			<div id="box2"> 
				<p class="Date">{{.Month }}</p>
				<p class="Date">{{.Day1 }}</p>
				<p class="Date">{{.Weekday }}</p>
			</div>
			<div id="box3"> 
				<pre><p class="time">{{.Hour}} : {{.Minute}}</p></pre>
			</div>
		</div>
	</div>
</body>
