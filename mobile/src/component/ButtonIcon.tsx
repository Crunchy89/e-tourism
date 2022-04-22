import React from 'react';
import {Image,Text,TouchableOpacity,Button} from 'react-native';
import {StyleSheet} from 'react-native';

const ButtonIcon = (props:any) => {
    return (
        <TouchableOpacity onPress={()=>{props.navigation.navigate(props.link)}} style={styles.button}>
            <Image source={props.image} style={{width:50,height:50}}/>
            <Text>{props.title}</Text>
        </TouchableOpacity>
    );
};

const styles = StyleSheet.create({
    button:{
        flexDirection:"column",
        alignItems:"center",
        width:80,
        height:80,
    },
    text:{
        textAlign:"justify"
    }
  });

export default ButtonIcon;