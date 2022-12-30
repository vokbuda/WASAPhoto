

export default class Comment {
    authorid=null
    postid=null
    commentid=null
    text=null
    quantityLikes=null
    quantityDislikes=null
    me=null
    lastChange=null
    currentemotion=null
    username=null
 
    constructor(authorid, postid,commentid,username,text, quantityLikes,quantityDislikes,me,lastChange,currentemotion) 
    {
      this.authorid = authorid
      this.postid = postid
      this.text=text
      this.quantityLikes=quantityLikes
      this.quantityDislikes=quantityDislikes
      this.me=me
     
      this.lastChange=lastChange
      this.currentemotion=currentemotion
      this.username=username
      this.commentid=commentid
    }
    
}
