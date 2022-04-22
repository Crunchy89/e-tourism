import React from 'react';
import {StyleSheet} from 'react-native';
import LinearGradient from 'react-native-linear-gradient';
import ButtonIcon from './ButtonIcon';
import { menu } from './data';


const Menu  = (props) => {
    return (
        <LinearGradient start={{x: 0, y: 0}} end={{x: 1, y: 0}} colors={['rgba(255,226,89,.8)', 'rgba(255,204,86,.8)', 'rgba(255,167,81,.8)']} style={styles.menuContainer}>
            {menu.map((item,index) => (
                <ButtonIcon 
                    key={index}
                    title={item.title}
                    image={item.image}
                    link={item.link}
                    navigation={props.navigation}
                />
            ))}
        </LinearGradient>
    );
};

const styles = StyleSheet.create({
    menuContainer:{
        flex:3,
        flexDirection:"row",
        justifyContent:"space-between",
        alignContent:"space-around",
        borderRadius:5,
        marginTop:10,
        marginBottom:10,
        flexWrap:"wrap",
        padding: 10,
    }
  });

export default Menu;