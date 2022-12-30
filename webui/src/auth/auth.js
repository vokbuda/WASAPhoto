


function isAuthenticated(){
    
    if (sessionStorage.getItem("token")==null){
        return  true
    }else{
        return false
    }

}
export default isAuthenticated