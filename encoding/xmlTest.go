package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
)

type rss2 struct {
	XMLName    xml.Name `xml:"rss"`
	ChannelDoc channel  `xml:"channel"`
}

type channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
	He          string   `xml:"hehe"` //这是不存在的节点
	Item        []item   `xml:"item"`
}

type item struct {
	XMLName xml.Name `xml:"item"`
	Title   string   `xml:"title"`
	Guid    guid     `xml:"guid"`
}

type guid struct {
	GuidValue string `xml:",chardata"`
	GuidAttr  string `xml:"isPermaLink,attr"`
}

func main() {

	s := `<?xml version="1.0" encoding="UTF-8" ?>
<rss version="2.0">
<channel>
 <title>RSS Title</title>
 <description>This is an example of an RSS feed</description>
 <link>http://www.example.com/main.html</link>
 <lastBuildDate>Mon, 06 Sep 2010 00:01:00 +0000 </lastBuildDate>
 <pubDate>Sun, 06 Sep 2009 16:20:00 +0000</pubDate>
 <ttl>1800</ttl>
 <item>
  <title>Example entry</title>
  <description>Here is some text containing an interesting description.</description>
  <link>http://www.example.com/blog/post/1</link>
  <guid isPermaLink="true">7bd204c6-1655-4c27-aeee-53f933c5395f</guid>
  <pubDate>Sun, 06 Sep 2009 16:20:00 +0000</pubDate>
 </item>
 <item>
  <title>Example entry2</title>
  <description>Here is some text containing an interesting description.</description>
  <link>http://www.example.com/blog/post/12</link>
  <guid isPermaLink="true">7bd204c6-1655-4c27-aeee-53f933c53952f</guid>
  <pubDate>Sun, 06 Sep 2009 16:20:00 +0000</pubDate>
 </item>
</channel>
</rss>`

	st := new(rss2)

	err := xml.NewDecoder(bytes.NewBuffer([]byte(s))).Decode(st)

	if err == nil {
		fmt.Println(st)
	} else {
		fmt.Println(err)
	}

	err = xml.NewEncoder(os.Stdout).Encode(st)

	fmt.Println()

	st2 := new(rss2)

	err = xml.Unmarshal([]byte(s), st2)
	if err == nil {
		fmt.Println(st2)
	}

	result, _ := xml.Marshal(st2)
	fmt.Println(string(result))
	fmt.Println(st2.ChannelDoc.Item[0].Guid.GuidAttr)
}
