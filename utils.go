package main

func getIDFromMention(content string) string {
  return content[3 : len(content)-1]
}
