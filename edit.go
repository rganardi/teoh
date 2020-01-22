package main

var (
	template = []byte(`
<html>
	<head>
<style>
body {
	margin:1em auto;
	width:100vw;
	max-width:40em;
	padding:0 .62em;
	font:1.2em/1.62 monospace;
}
@media print {
	body {
		max-width:none
	}
}
textarea {
	display:block;
	padding:1em;
	border:1px solid #aaa;
	resize:none;
	margin:0;
}
button {
	display:inline;
	padding:0.4em;
	margin:0.2em 0;
	height:2em;
	background:none;
	border:1px solid #aaa;
	color:black;
}
textarea, button {
	box-sizing:border-box;
	border-radius:0.3em;
	transition:0.2s;
	font-family:inherit;
	font-size:100%;
}
textarea:hover, button:hover {
	border-color:#000;
}
#status {
	font-size:0.8em;
	margin:0.2em 0.4em;
	text-align:right;
}
#container {
	display:flex;
	flex-direction:column;
	height:80vh;
	width:100%;
}

</style>
	</head>
	<body>
		<div id='container'>
		<textarea id='file' cols='72' rows='20'></textarea>
		<span id='status'></span>
		<br>
		<button id='refresh' onclick='refresh()'>
			Refresh
		</button>
		<button id='submit' onclick='submit()'>
			Send
		</button>
		</div>
	<script>
		const textarea = document.getElementById('file')
		const status = document.getElementById('status')
		const controller = new AbortController();

		const get = async () => {
			status.innerText = 'fetching'
			try {
				const response = await fetch('/');
				if (!response.ok) {
					throw new Error(response.statusText);
				}
				status.innerText = 'done'
				setTimeout(() => {
					if (status.innerText === 'done') {
						status.innerText = ''
					}
				}, 2000);

				return await response.text();
			} catch (err) {
				console.error(err);
				status.innerText = 'error fetching file'
				return '';
			}
		}

		const refresh = async () => {
			textarea.value = await get();
		}

		const submit = async () => {
			const body = textarea.value

			if (status.innerText === 'sending') {
				controller.abort();
			}

			status.innerText = 'sending'

			try {
				const response = await fetch('/', {
					method: 'POST',
					headers: {
						'Content-Type': 'text/plain',
					},
					body,
				}, { signal: controller.signal })
				if (!response.ok) {
					throw new Error(response.statusText);
				}

				textarea.value = await response.text();
				status.innerText = 'done'

				setTimeout(() => {
					if (status.innerText === 'done') {
						status.innerText = ''
					}
				}, 2000);
			} catch (err) {
				console.error(err);
				status.innerText = 'error sending file'
			}
		}

		refresh();

	</script>
	</body>
</html>
`)
)
