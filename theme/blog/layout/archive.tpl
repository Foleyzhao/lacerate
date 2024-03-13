{{ define "content" }}
    {{ range .archive }}
    <div class="archive animated fadeInDown">
        <div class="list-with-title">
            <div class="listing-title">{{ .YearStr }}</div>
            <ul class="listing">
                {{ range .Months }}
                <div class="listing-item">
                    {{ range .Posts }}
                    <div class="listing-post"><a href="{{ .Url }}" title="{{ .Title }}">{{ .Title }}</a>
                        <div class="post-time"><span class="date">{{ .CreatedAt | format }}</span>
                        </div>
                    </div>
                    {{ end }}
                </div>
                {{ end }}
            </ul>
        </div>
    </div>
    {{ end }}
{{ end }}