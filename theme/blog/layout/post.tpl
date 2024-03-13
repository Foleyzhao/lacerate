{{ define "content" }}
<div class="post-page">
    <div class="post animated fadeInDown">
        <div class="post-title">
            <h3>{{ .post.Title }}
            </h3>
        </div>
        <div class="post-content" id="content">
            {{ .post.Content | unescaped }}
            <nav class="article-nav" id="state">
                <span class="label label-important">PERMANENT LINK:</span>
                <a href="https://happy.zj.cn{{.post.Url}}">https://happy.zj.cn{{.post.Url}}</a>
            </nav>
        </div>
        <div class="post-footer">
            <div class="meta">
                <div class="info">
                    <i class="fa fa-calendar"></i>
                    <span class="date">{{ .post.CreatedAt | format }}</span>
                    &nbsp;&nbsp;
                    <i class="fa fa-folder-open"></i>
                    {{ range .post.Category }}
                    <a href="/category/{{ . }}">{{ . }}</a>&nbsp;
                    {{ end }}
                    &nbsp;&nbsp;
                    <i class="fa fa-tags"></i>
                    {{ range .post.Tags }}
                    <a href="/tag/{{ . }}">{{ . }}</a>&nbsp;
                    {{ end }}
                </div>
            </div>
        </div>
    </div>
    <!--评论-->
</div>
{{ end }}