
$(document).ready(function(){
	$('a[href*="/gfw.go101.org/"]').each(function () {
		let p = "/gfw.go101.org"
		let h = $(this).attr("href")
		let i = h.indexOf(p)
		if (i >= 0) {
			$(this).attr("href", h.substr(i+p.length));
		}
	});
});

