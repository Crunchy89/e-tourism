import React from 'react';
import {TouchableOpacity,View,Text,StyleSheet,Image} from "react-native"
import ArrowBack from "../assets/img/arrowBack.png"

const ButtonBack = (props:any) => {
    return (
        <TouchableOpacity onPress={()=>props.navigation.navigate(props.link)} style={style.button}>
            <View style={style.container}>
                <Image source={ArrowBack} />
                <Text style={style.text}>{props.title}</Text>
            </View>
        </TouchableOpacity>
    );
};

const style = StyleSheet.create({
    container:{
        flexDirection:"row",
        alignItems:"center"
    },
    button:{
        margin:10
    },
    text:{
        fontSize:20,
        fontWeight:"bold",
        color:"#fff",
        marginLeft:10
    }
})

export default ButtonBack;