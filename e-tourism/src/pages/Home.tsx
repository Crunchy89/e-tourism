import React from 'react';

import {SafeAreaView, View,Image,StyleSheet} from 'react-native';
import logo from '../assets/img/logo.png';
import News from "../component/News"
import Menu from "../component/Menu"
const Home = ({navigation}) => {
    
    return (
        <SafeAreaView style={styles.container}>
            <View style={styles.logo}>
                <Image source={logo} />
            </View>
            <News/>
            <Menu navigation={navigation} />
        </SafeAreaView>
    );
};

const styles = StyleSheet.create({
    container: {
      flex: 3,
      padding: 20,
      justifyContent: 'center',
    },
    logo:{
        marginBottom:5,
    },
  });

export default Home;