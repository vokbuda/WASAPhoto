export default class Post {
    constructor(authorid, postid,text, quantityLikes,quantityDislikes,me,image,lastChange,currentemotion) 
    {
      this.authorid = authorid
      this.postid = postid
      this.text=text
      this.quantityLikes=quantityLikes
      this.quantityDislikes=quantityDislikes
      this.me=me
      this.image=image
      this.lastChange=lastChange
      this.currentemotion=currentemotion
    }
    
}