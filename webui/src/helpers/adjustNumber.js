function adjustNumber(num_entities){
    
    if (num_entities >= 1000000) {
        return Math.round(parseFloat(num_entities/1000000.0)*100)/100 + "m"

       
    }
    if(num_entities >= 1000) {
        return Math.round(parseFloat(num_entities/1000.0)*100)/100 + "k"
    }
    
    return num_entities.toString()
}
export {adjustNumber}

