javascript:(function() {
	try {
		var d = document,
		    w =
		    window;
		if (!d.body || d.body.innerHTML == '')
			throw (0);
		var s = d.createElement('link'),
		    h = d.getElementsByTagName('head')[0],
		    i = d.createElement('div'),
		    j = d.createElement('script');
		s.rel = 'stylesheet';
		s.href = '//dotepub.com/s/dotEPUB-favlet.css';
		s.type = 'text/css';
		s.media = 'screen';
		h.appendChild(s);
		i.setAttribute('id', 'dotepub');
		i.innerHTML = '<div id="status">	<p>		Conversion in progress...	</p></div>';
		d.body.appendChild(i);
		j.type = 'text/javascript';
		j.charset = 'utf-8';
		j.src = '//dotepub.com/j/dotepub.js?v=1.2&s=ask&t=mobi&g=en';
		h.appendChild(j);
	} catch(e) {
		w.alert('The page has no content or it is not fully loaded. Please, wait till the page is loaded.');
	}
})(); 