function convertToString(num_entities,changed){
   
    
    
    if(parseInt(num_entities,10).toString()===num_entities) {
        return (parseInt(num_entities,10)+changed).toString()
      }

      
    return num_entities

    
}

export {convertToString}