# github-rss

github public activity to rss feed.

## Usage
* `github-rss` gets public user activity received events (TODO: only returns starred, created and pushed repos, returning a description of each repo would involve too many request and a rate limit would be hit soon enough)
* `gitorg-rss` gets a list of repos created by an organization
``` shell
> github-rss azimut
<?xml version="1.0" encoding="UTF-8"?><feed xmlns="http://www.w3.org/2005/Atom">
  <title>azimut github activity</title>
  <id>https://github.com/azimut</id>
  <updated>2021-01-24T16:49:28-03:00</updated>
  <subtitle>discussion about tech, footie, photos</subtitle>
  <link href="https://github.com/azimut"></link>
  <entry>
    <title>asears starred MaterialDesignInXAML/MaterialDesignInXamlToolkit</title>
    <updated>2021-01-24T18:02:16Z</updated>
    <id>tag:github.com,2021-01-24:/MaterialDesignInXAML/MaterialDesignInXamlToolkit</id>
    <link href="https://github.com/MaterialDesignInXAML/MaterialDesignInXamlToolkit" rel="alternate"></link>
    <summary type="html"></summary>
  </entry>
  <entry>
    <title>okbel starred DrBoolean/immutable-ext</title>
    <updated>2021-01-24T18:00:37Z</updated>
    <id>tag:github.com,2021-01-24:/DrBoolean/immutable-ext</id>
    <link href="https://github.com/DrBoolean/immutable-ext" rel="alternate"></link>
    <summary type="html"></summary>
  </entry>
  <entry>
    <title>soredake starred mystor/micro-coreutils</title>
    <updated>2021-01-24T17:53:58Z</updated>
    <id>tag:github.com,2021-01-24:/mystor/micro-coreutils</id>
    <link href="https://github.com/mystor/micro-coreutils" rel="alternate"></link>
    <summary type="html"></summary>
  </entry>
  <entry>
    <title>soredake starred frostming/pdm</title>
    <updated>2021-01-24T17:41:09Z</updated>
    <id>tag:github.com,2021-01-24:/frostming/pdm</id>
    <link href="https://github.com/frostming/pdm" rel="alternate"></link>
    <summary type="html"></summary>
  </entry>
  <entry>
    <title>katychuang starred hahey/lanyon-hakyll</title>
    <updated>2021-01-24T17:36:53Z</updated>
    <id>tag:github.com,2021-01-24:/hahey/lanyon-hakyll</id>
    <link href="https://github.com/hahey/lanyon-hakyll" rel="alternate"></link>
    <summary type="html"></summary>
  </entry>
  <entry>
    <title>jedisct1 starred microsoft/win32metadata</title>
    <updated>2021-01-24T16:20:34Z</updated>
    <id>tag:github.com,2021-01-24:/microsoft/win32metadata</id>
    <link href="https://github.com/microsoft/win32metadata" rel="alternate"></link>
    <summary type="html"></summary>
  </entry>
  <entry>
    <title>soredake starred felipefialho/frontend-challenges</title>
    <updated>2021-01-24T16:10:56Z</updated>
    <id>tag:github.com,2021-01-24:/felipefialho/frontend-challenges</id>
    <link href="https://github.com/felipefialho/frontend-challenges" rel="alternate"></link>
    <summary type="html"></summary>
  </entry>
  <entry>
    <title>soredake starred Asabeneh/30-Days-Of-React</title>
    <updated>2021-01-24T16:10:43Z</updated>
    <id>tag:github.com,2021-01-24:/Asabeneh/30-Days-Of-React</id>
    <link href="https://github.com/Asabeneh/30-Days-Of-React" rel="alternate"></link>
    <summary type="html"></summary>
  </entry>
  <entry>
    <title>ryukinix starred Lissy93/personal-security-checklist</title>
    <updated>2021-01-24T15:24:40Z</updated>
    <id>tag:github.com,2021-01-24:/Lissy93/personal-security-checklist</id>
    <link href="https://github.com/Lissy93/personal-security-checklist" rel="alternate"></link>
    <summary type="html"></summary>
  </entry>
  <entry>
    <title>soredake starred NateGraff/blessedvirginmary</title>
    <updated>2021-01-24T14:57:12Z</updated>
    <id>tag:github.com,2021-01-24:/NateGraff/blessedvirginmary</id>
    <link href="https://github.com/NateGraff/blessedvirginmary" rel="alternate"></link>
    <summary type="html"></summary>
  </entry>
  <entry>
    <title>soredake starred veeral-patel/how-to-secure-anything</title>
    <updated>2021-01-24T14:55:09Z</updated>
    <id>tag:github.com,2021-01-24:/veeral-patel/how-to-secure-anything</id>
    <link href="https://github.com/veeral-patel/how-to-secure-anything" rel="alternate"></link>
    <summary type="html"></summary>
  </entry>
  <entry>
    <title>soredake starred HackerPolice/MissedIT</title>
    <updated>2021-01-24T14:54:50Z</updated>
    <id>tag:github.com,2021-01-24:/HackerPolice/MissedIT</id>
    <link href="https://github.com/HackerPolice/MissedIT" rel="alternate"></link>
    <summary type="html"></summary>
  </entry>
  <entry>
    <title>soredake starred itsme2417/EyeHook</title>
    <updated>2021-01-24T14:54:40Z</updated>
    <id>tag:github.com,2021-01-24:/itsme2417/EyeHook</id>
    <link href="https://github.com/itsme2417/EyeHook" rel="alternate"></link>
    <summary type="html"></summary>
  </entry>
  <entry>
    <title>soredake starred abhisheknaiidu/awesome-github-profile-readme</title>
    <updated>2021-01-24T14:53:40Z</updated>
    <id>tag:github.com,2021-01-24:/abhisheknaiidu/awesome-github-profile-readme</id>
    <link href="https://github.com/abhisheknaiidu/awesome-github-profile-readme" rel="alternate"></link>
    <summary type="html"></summary>
  </entry>
  <entry>
    <title>diasbruno starred mafintosh/autoname</title>
    <updated>2021-01-24T14:29:55Z</updated>
    <id>tag:github.com,2021-01-24:/mafintosh/autoname</id>
    <link href="https://github.com/mafintosh/autoname" rel="alternate"></link>
    <summary type="html"></summary>
  </entry>
  <entry>
    <title>patriciogonzalezvivo starred MIT-SPARK/Kimera</title>
    <updated>2021-01-24T14:27:53Z</updated>
    <id>tag:github.com,2021-01-24:/MIT-SPARK/Kimera</id>
    <link href="https://github.com/MIT-SPARK/Kimera" rel="alternate"></link>
    <summary type="html"></summary>
  </entry>
  <entry>
    <title>diasbruno starred yitzchak/common-lisp-jupyter</title>
    <updated>2021-01-24T14:27:37Z</updated>
    <id>tag:github.com,2021-01-24:/yitzchak/common-lisp-jupyter</id>
    <link href="https://github.com/yitzchak/common-lisp-jupyter" rel="alternate"></link>
    <summary type="html"></summary>
  </entry>
</feed>
```
~
