/**
 * Sample React Native App
 * https://github.com/facebook/react-native
 *
 * @format
 * @flow strict-local
 */

import React from 'react';
import {SafeAreaView,ImageBackground} from "react-native"
import bg from "./src/assets/img/bg.png"
import Home from "./src/pages/Home";
import Destination from './src/pages/Destination';
import Market from './src/pages/Market';
import Hotel from "./src/pages/Hotel"
import Flight from "./src/pages/Flight"
import Local from "./src/pages/Local"
import Rental from "./src/pages/Rental"
import Food from "./src/pages/Food"
import Hospital from "./src/pages/Hospital"
import Rest from "./src/pages/Rest"
import Police from "./src/pages/Police"
import Event from "./src/pages/Event"
import Help from "./src/pages/Help"

import { NavigationContainer,DefaultTheme} from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';
import {LogBox} from "react-native";


const MyTheme = {
  ...DefaultTheme,
  colors: {
    ...DefaultTheme.colors,
    background: 'transparent',
  },
};

const Stack = createNativeStackNavigator();
const App = () => {
  LogBox.ignoreLogs([
    "exported from 'deprecated-react-native-prop-types'.",
    ])
  return (
    <SafeAreaView>
      <ImageBackground source={bg} style={{width: '100%', height: '100%'}}>
      <NavigationContainer theme={MyTheme}>
          <Stack.Navigator 
          initialRouteName="Home" 
          detachPreviousScreen={true} 
          detachInactiveScreens={true}
          >
            <Stack.Screen  name="Home" component={Home} options={{headerShown: false}} />
            <Stack.Screen  name="Destination" component={Destination} options={{headerShown: false}}  />
            <Stack.Screen  name="Market" component={Market} options={{headerShown: false}}  />
            <Stack.Screen  name="Hotel" component={Hotel} options={{headerShown: false}}  />
            <Stack.Screen  name="Flight" component={Flight} options={{headerShown: false}}  />
            <Stack.Screen  name="Local" component={Local} options={{headerShown: false}}  />
            <Stack.Screen  name="Rental" component={Rental} options={{headerShown: false}}  />
            <Stack.Screen  name="Food" component={Food} options={{headerShown: false}}  />
            <Stack.Screen  name="Hospital" component={Hospital} options={{headerShown: false}}  />
            <Stack.Screen  name="Rest" component={Rest} options={{headerShown: false}}  />
            <Stack.Screen  name="Police" component={Police} options={{headerShown: false}}  />
            <Stack.Screen  name="Event" component={Event} options={{headerShown: false}}  />
            <Stack.Screen  name="Help" component={Help} options={{headerShown: false}}  />
          </Stack.Navigator>
        </NavigationContainer>
      </ImageBackground>
    </SafeAreaView>
  );
};

export default App;
