import React from 'react';
import {View,Text} from 'react-native';
import Container from '../component/Container';
import ButtonBack from '../component/ButtonBack';

const Help = ({navigation}) => {
    return (
        <>
        <ButtonBack link="Home" navigation={navigation} title="Help"/>
        <Container>
            <View>
                <Text>Help</Text>
            </View>
        </Container>
        </>
    );
};

export default Help;