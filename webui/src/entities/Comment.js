

export default class Comment { 
    constructor(authorid, postid,commentid,username,text, quantityLikes,quantityDislikes,me,lastChange,currentemotion, avatar) 
    {
      this.authorid = authorid
      this.postid = postid
      this.text=text
      this.quantityLikes=quantityLikes
      this.quantityDislikes=quantityDislikes
      this.me=me
      this.avatar=avatar
      this.lastChange=lastChange
      this.currentemotion=currentemotion
      this.username=username
      this.commentid=commentid
    }
    
}
