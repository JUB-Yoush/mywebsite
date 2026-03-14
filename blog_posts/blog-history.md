---
title: 'All 7 Versions of my blog'
date: '2026-03-14'
tags: ['reflection','web','programming']
---
[Fraud was sick play the update](https://youtu.be/B6ltSXD9eMU?si=uDPtXN3Y5XAThYBa)

# History of my Personal Sites
I remade my blog again so here is a history of every version of my blog that has ever existed.
## Ghost, 10/2021
My first website was made using [Ghost](https://ghost.org/) simply because [Ali Abdaal](<https://youtu.be/acBJsjCqgtM?si=vkAEhL0LO6hMxXJc>) told me to use it and he was the authority I was listening to at the time. He only mentions the price of the hosted solution and doesn't mention you can host it yourself at all (he does mention that it's open source). 

Throughout the entire video He never mentions a static generator, with Ghost being the recommendation for people who "know a little bit about computers". I was already familiar with HTML back in 2021, but I thought making a website in HTML was like writing a native application in x86 assembly, something that was impractical and abstracted away by modern tools that "real" developers use. I did not know markdown nor HTML templates existed so I was under the impression that the only ways to author content for the web was writing articles in a webpage with a cms based system like WordPress/Ghost, or just writing HTML by hand. 

I didn't end up sticking with the Ghost site for very long at all. I think I wrote my Drawabox blog post (which I [recently archived](https://jaydenpb.netblog/2022/01/21/yb1-draw-the-box-now-do-it-249-more-times/) there first and then re-wrote it in WordPress but I cant remember. I remember I found the process of configuring a [DigitalOcean](https://www.digitalocean.com/) droplet to host the site kind of confusing.
## WordPress, 01/2022
I decided to go with a [WordPress](https://wordpress.org/) site over a Ghost one, simply because it was easier to set up in DigitalOcean and I found a theme that I preferred to anything ghost offered (I remember at the time Ghost only had like, 2?). 

I wrote my first two blog posts: one on Drawabox and one on Godot (which I also [recently archived](https://jaydenpb.net/blog/2022/03/22/yb2-i-used-godot-for-three-years-and-learned-nothing-here-is-how-to-avoid-that/)).  
## Raw html, 03/2022
I was kind of in the peak of my "Emulating everything [James Scholz](https://www.youtube.com/@jvscholz) does" era, which actually did help make large sweeping valuable changes to my life. Of all the role models I could've found at the time, he was a really good one. I don't watch him much nowadays (to be fair, don't watch much of anyone on YouTube anymore) as I feel like I was too caught up in feelings of comparison, of all of me compared to the version of himself who I had extrapolated based on what he presented of himself publicly. Wasn't very healthy for me. 

Anyway part of that emulation involved his website which I blatantly stole. His site still lives [here](https://jvscholz.com/) although it seems outdated. I could only assume he simply wrote HTML files directly and that's what I did as well. It worked fine, I had a blog page set up but didn't write anything, as my old posts were written in the WordPress cms.
## Frozen flask site, 01/2023
Baby's first generated static site! My motivation was mostly wanting a way to make blog posts within a static site without writing html by hand. I had browsed a bit around with static site generators like [Hugo](https://gohugo.io/) and [Jekyll](https://jekyllrb.com/) but didn't like how they were theme based, I wanted to design the website myself with HTML templates (but I didn't know what those were at the time).

I was familiar with [Flask](https://flask.palletsprojects.com/en/stable/) through it's use at a hackathon I had participated in not long before. It taught me about [Jinja](https://jinja.palletsprojects.com/en/stable/templates/) HTML templates. Exactly what I was looking for! I ended up making a site using the [flask-freeze](https://frozen-flask.readthedocs.io/en/stable/index.html) package that would walk through all your api endpoints and pre-render a HTML page for each one. I had followed a tutorial to set it up but I can no longer find it unfortunately.

The web design of this version of the site is just about the same as the current, with some slight alterations. This design was also blatantly stolen from Quinn Ha's previous personal site, but he's since abandoned it for a different much more [inspired design](https://www.quinnha.menu/), I had stolen his site on a whim after discovering his LinkedIn page, and in the smallest of worlds coincidences, one of his friends ended up interning with me at Scotiabank and asked me why my site looked so much like Quinns. I was then reached out to Quinn and he gave he his blessing for my plagiarism, he also helped with preparing for my Google Internship, as he was there the summer before me.
## SvelteKit + Vercel, 03/2023
I'm something of a contrarian. Or I'm a fan of picking the best tools rather, even if they aren't popular (which they often aren't). Sometimes the popular things end up being the best, like when I switched from Emacs Org mode to Obsidian (although .org is still great and preferable to a certain kind of hacker for sure). Anyway when it finally came time to learn about JavaScript frameworks I gravitated towards [Svelte](https://svelte.dev/) over the ever ubiquitous [React](https://react.dev/) simply because people on the internet said it was better. This is just about the same reason I ended up learning Godot instead of Unity. 

Overall I think my propensity to lean towards niche tools has been a net good, as it's allowed me to gain perspective working with a wider range of tools while still being forced to interact with popular defaults as they're popular and the default. React has an absolute monopoly on front-end JavaScript frameworks despite only really being the best at being popular, so from solely a career perspective it doesn't really make much sense to learn anything else unfortunately.

I re-wrote the flask site as a server-less [SvelteKit](https://svelte.dev/docs/kit/introduction) application and hosted it for free on [Vercel](https://vercel.com). I mostly followed the [excellent blog post](https://joshcollinsworth.com/blog/build-static-sveltekit-markdown-blog) From Josh Collinsworth, who also recently happened to hit the nail on the head on exactly how I feel about the current [ai slop world we find ourselves in](https://joshcollinsworth.com/blog/sloptimism). Also consider reading what [Bob Nystrom wrote](https://www.journal.stuffwithstuff.com/2026/01/24/the-value-of-things/) for a different angle on the topic I also really like.
## SvelteKit-Static + gh-pages, 04/2024
Eventually I had learned about static hosting and [Github Pages](https://docs.github.com/en/pages). Vercel was cool and free but I didn't really understand how "serverless" architecture worked (I still don't really), and idea of having pre-rendered html pages that github rendered for me made sense. Not having to concern myself with two separate platforms was also great.

When I was remaking the GDYU website I had used SvelteKit and modeled it after what I learned from making my site, I did start with static exporting from the jump, and after finishing that, retroactively remade my site with static rendering. This is the version of the site that I had used for the longest.
## Go + Templ, 02/2026
And now we've gone full circle.

Going from HTML templates to dynamic full-stack frontend framework in vercel to statically generated front end framework site back to HTML templates feels like I've finally reached the enlightenment end of the Soyjack bell-curve meme.

SvelteKit is designed for dynamic server based production web applications. Not static blogs like mine. That's not to say you can't make simple static blog in SvelteKit, and of all the metaframeworks I've used it's very beginner friendly and grows with your understanding. But using a JavaScript metaframework to parse some markdown and render some html is like using a shotgun to kill a housefly. Looking over my repository there are a bunch of configuration files and 62.7MB of node modules who's purpose in the creation of the site I couldn't communicate to you. Not to say it was super complex, if I reviewed modern JS stuff and read over the code again I could probably figured it out, but I remembered how my site used to just be writing html by hand, how complicated should making a my site be?

I think one of the best things that switching to [Raylib](https://www.raylib.com/index.html) for game development has taught me is that a lot of the complexity we're used to working within is optional. I think a lot of "best practices" and popular tools/engines lead to complicated software, as they're designed with the patterns and solution that assume you need the complexity. 

I don't think SvelteKit or Unity should be simpler, When you're programming something that takes advantage of the broad toolset they provide it can be a tremendous asset. But when I consider a project's technology I now tend to pick tools based on my specific needs, not just opt in for the big "do everything" tool people tend to pick. And I've found that I often don't need a lot of the stuff the big tools tend to give you. I'm quite happy writing simple code, that solves my simple problems, with minimal dependencies.

With this, I set out to remake my website with technology I could understand.

My new site is created with [Go](https://go.dev/) + [Templ templates](https://templ.guide/) (I could get away with using the regular html Go templates, but I like Templ files). 
The "site generator" is one ``main.go`` file that does the following:
- create a folder called /public.
- copy over the static folder that contains css/js/images into /public.
- create subdirectories for each route within /public.
- write the index.html page for each static pages using templ templates.
- parse my folder of markdown posts into structs containing their html + metadata with [goldmark](https://github.com/yuin/goldmark).
- generate a directory for the post + write it's content.
- make a page filtering posts by every tag.
- write an rss file.

And all I have to do is serve the static site in public. The deployment file for github pages automatically builds the site and hosts the content with /public being the root. 

The only major features I'd like but haven't implemented yet would be having foonotes displayed inline to the left/below paragraphs but that would require me to parse the AST of goldmark (which isn't that bad, I just haven't gotten to it yet.), But I'm very happy with the current state of my blog and think this style (bespoke ssg that copies and writes files directly) is what I'll stick with.

End of blog post CTA: If you have a super cool blog (and are also a game developer in the Toronto Area) then consider including your site in the new [Toronto Gamedev Webring](https://webring.jaydenpb.net/) I put together. Please! It can't just be me and Alvin in there forever!
