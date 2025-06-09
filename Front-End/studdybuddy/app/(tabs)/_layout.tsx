import { Tabs } from 'expo-router';
import React from 'react';
import { Platform } from 'react-native';

import { HapticTab } from '@/components/HapticTab';
import { IconSymbol } from '@/components/ui/IconSymbol';
import TabBarBackground from '@/components/ui/TabBarBackground';
import { Colors } from '@/constants/Colors';
import { useColorScheme } from '@/hooks/useColorScheme';

export default function TabLayout() {
  const colorScheme = useColorScheme();

  return (
    <Tabs
      screenOptions={{
        tabBarActiveTintColor: "#ffffff", // Texto ativo branco
        tabBarInactiveTintColor: "#ffffff", // Texto inativo branco
        headerShown: false,
        tabBarButton: HapticTab,
        tabBarStyle: Platform.select({
          ios: {
            backgroundColor: "#1F2937",
            position: 'absolute',
            borderTopWidth: 0,
          },
          default: {
            backgroundColor: "#1F2937",
            flexDirection: "row", // Isso já é o padrão para tab bars
          }
        }),
        tabBarItemStyle: {
          flexDirection: "column", // Organiza ícone e texto verticalmente
          justifyContent: "center",
          alignItems: "center",
        },
      }}
    >
      <Tabs.Screen
        name="index"
        options={{
          title: 'Home',
          tabBarIcon: () => <IconSymbol size={28} name="house.fill" color={"#ffffff"} />,
        }}
      />
      <Tabs.Screen
        name="explore"
        options={{
          title: 'Explore',
          tabBarIcon: () => <IconSymbol size={28} name="paperplane.fill" color={"#ffffff"} />,
        }}
      />
    </Tabs>
  );
}