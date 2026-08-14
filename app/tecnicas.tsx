import { View, Text, Pressable } from 'react-native';
import { ScreenContent } from 'components/screen/ScreenContent';
import MaterialCommunityIcons from '@expo/vector-icons/MaterialCommunityIcons';
import FontAwesome5 from '@expo/vector-icons/FontAwesome5';
import Entypo from '@expo/vector-icons/Entypo';
import CardTechnique from 'components/ui/CardTechnique';
import { useRouter } from 'expo-router';

export default function Tecnicas() {
    const router = useRouter();

    return (
        <ScreenContent title="Técnicas" path="/tecnicas">
            {/* Titulo da página */}
            <Text className="text-2xl font-bold text-gray-800 dark:text-white">
                Técnicas de Estudo
            </Text>
            <View>
                <CardTechnique title="Pomodoro"    description="Tempo de estudo + Tempo de descanso"         cor="red"    paddingX="3"              icon={<MaterialCommunityIcons name="timer" size={28} color="#dc2626" />}       redirect={() => router.push('/pomodoro')} />
                <CardTechnique title="Feynman"     description="Ensine o conteúdo para reforçar aprendizado" cor="yellow" paddingX="2"              icon={<FontAwesome5 name="user-friends" size={24} color="#ca8a04" />}          redirect={() => router.push('/feynman')}  />
                <CardTechnique title="Mapa Mental" description="Organize ideias visualmente"                 cor="blue"   paddingX="3" paddingY="8" icon={<Entypo name="blackboard" size={24} color="#2563eb" />}                  redirect={() => router.push('/mindMap')}  />
                <CardTechnique title="Simulados"   description="Pratique com questões de provas anteriores"  cor="green"  paddingX="3"              icon={<MaterialCommunityIcons name="list-status" size={28} color="#16a34a" />} redirect={() => router.push('/cornell')}  />
            </View>
        </ScreenContent>
    );

}