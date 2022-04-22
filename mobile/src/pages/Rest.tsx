import React from 'react';
import {View,Text} from 'react-native';
import Container from '../component/Container';
import ButtonBack from '../component/ButtonBack';

const Rest = ({navigation}) => {
    return (
        <>
        <ButtonBack link="Home" navigation={navigation} title="Rest Area"/>
        <Container>
            <View>
                <Text>Rest</Text>
            </View>
        </Container>
        </>
    );
};

export default Rest;