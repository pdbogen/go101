index.html: index.html.pre index.svg index.html.post
	cat index.html.pre > index.html
	cat index.svg >> index.html
	cat index.html.post >> index.html

index.svg: index.dot
	dot -Tsvg index.dot > index.svg
