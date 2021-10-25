fun main (){
    val x : MutableList<String> = mutableListOf("salam", "alo", "alo")
    x.add("3")
    for ((i,v) in x.withIndex()){
        if (i == 2){
            x.remove(v)
        }
    }
    for ((i,v) in x.withIndex()){
        println("$i : $v")
    }

}
