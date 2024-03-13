{{ define "content" }}
    {{ range .postList }}
    <div class="post animated fadeInDown">
        <div class="post-title">
            <h3><a style="color: #0077FF" href="{{ .Url }}">{{ .Title }}</a>
            </h3>
        </div>
        <div class="post-content">
            <div class="p_part">
                {{ .Summary | unescaped }}
            </div>
        </div>
        <div class="post-footer">
            <div class="meta">
                <div class="info">
                    <i class="fa fa-calendar"></i>
                    <span class="date">{{ .CreatedAt | format }}</span>
                    &nbsp;&nbsp;
                    <i class="fa fa-folder-open"></i>
                    {{ range .Category }}
                    <a href="/category/{{ . }}">{{ . }}</a>&nbsp;
                    {{ end }}
                    &nbsp;&nbsp;
                    <i class="fa fa-tags"></i>
                    {{ range .Tags }}
                    <a href="/tag/{{ . }}">{{ . }}</a>&nbsp;
                    {{ end }}
                </div>
            </div>
        </div>
    </div>
    {{ end }}
{{ end }}