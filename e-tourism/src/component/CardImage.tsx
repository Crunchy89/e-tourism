import React from 'react';
import {View,Image,Text, StyleSheet,TouchableOpacity} from 'react-native';


const CardImage = (props:any) => {
    return (
        <TouchableOpacity onPress={()=>{props.navigation.navigate(props.link,{title:props.title,id:props.id})}}>
        <View style={styles.container}>
            <Image source={{uri:props.image}} style={styles.image} />
            <Text>{props.title}</Text>
        </View>
        </TouchableOpacity>
    );
};

const styles =StyleSheet.create({
    container:{
        margin:2,
        backgroundColor:'#fff',
        width:145,
        height:150,
        padding:5
    },
    image:{
        width:"100%",
        height:100,
        borderRadius:3
    },
});

export default CardImage;