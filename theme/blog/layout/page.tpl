{{ define "content" }}
<div class="archive animated fadeInDown">
    <div class="list-with-title">
        <div class="listing-title">{{ .pageTitle }}({{ .count }})</div>
        <ul class="listing">
            {{ range .content }}
            <div class="listing-item">
                <div class="listing-post"><a href="{{ .Url }}" title="{{ .Title }}">{{ .Title }}</a>
                    <div class="post-time">
                        <span class="date">{{ .CreatedAt | format }}</span>
                    </div>
                </div>
            </div>
            {{ end }}
        </ul>
    </div>
</div>
{{ end }}
