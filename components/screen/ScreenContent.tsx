import FontAwesome6 from '@expo/vector-icons/FontAwesome6';
import Ionicons from '@expo/vector-icons/Ionicons';
import MaterialIcons from '@expo/vector-icons/MaterialIcons';
import SlideModalComponent from 'components/ui/SlideModal';
import { useToggleTheme } from 'hooks/useToggleTheme';
import { Button } from 'components/ui/Button';
import { useRouter } from 'expo-router';
import React, { type ComponentProps, useMemo, useState, useEffect } from 'react';
import { Pressable, ScrollView, Text, View } from 'react-native';
import { SafeAreaProvider, SafeAreaView } from 'react-native-safe-area-context';
import { StatusBar } from 'expo-status-bar';
import NotificationCard from 'components/ui/NotificationCard';



type TabId = 'home' | 'tecnicas' | 'ia' | 'focus' | 'config';

type BottomTab = {
  id: TabId;
  label: string;
  icon: ComponentProps<typeof FontAwesome6>['name'];
};

interface ScreenContentProps {
  title: string;
  path: string;
  children?: React.ReactNode;
  activeTab?: TabId;
  notificationCount?: number;
  isDark?: boolean;
  onHomePress?: () => void;
  onNotificationsPress?: () => void;
  onTabChange?: (tabId: TabId) => void;
  onThemeToggle?: () => void;
}

const bottomTabs: BottomTab[] = [
  { id: 'home', label: 'Início', icon: 'house' },
  { id: 'tecnicas', label: 'Técnicas', icon: 'lightbulb' },
  { id: 'ia', label: 'IA', icon: 'robot' },
  { id: 'focus', label: 'Foco', icon: 'bolt' },
  { id: 'config', label: 'Configurações', icon: 'gear' },
];

export const ScreenContent: React.FC<ScreenContentProps> = ({
  title,
  path,
  children,
  activeTab,
  notificationCount = 0,
  onHomePress,
  onNotificationsPress,
  onTabChange,
  onThemeToggle,
}) => {
  const selectedTab = useMemo(
    () => activeTab ?? getTabFromRoute(path, title),
    [activeTab, path, title]
  );
  const { isDark, setDarkTheme, setLightTheme } = useToggleTheme();
  
  const backgroundClassName = isDark ? 'flex-1 bg-gray-800' : 'flex-1 bg-white';
  const headerClassName = isDark
  ? 'border-b border-gray-700 bg-gray-800 shadow-sm'
  : 'border-b border-gray-200 bg-white shadow-sm';
  const textClassName = isDark ? 'text-gray-100' : 'text-gray-800';
  const mutedTextClassName = isDark ? 'text-gray-400' : 'text-gray-500';
  const iconColor = isDark ? '#f3f4f6' : '#1f2937';
  
  const [notificationModalVisible, setNotificationModalVisible] = useState(false);
  
  const router = useRouter();
  
  // temporary cariable remove after implementing the notification hook
  const mockNotifications = [
    {
      id: 1,
      title: 'Estudar React Native',
      time: '10:00',
      description: 'Revisar componentes e navegação',
    },
    {
      id: 2,
      title: 'Prova de Banco de Dados',
      time: '14:30',
      description: 'Estudar SQL e PostgreSQL',
    },
    {
      id: 3,
      title: 'Reunião do projeto',
      time: '18:00',
      description: 'Reunião com a equipe do Tripfy',
    },
  ];
  
  
  const handleNotificationModal = () => {
    setNotificationModalVisible(true);
  };

  return (
    <SafeAreaProvider>
      <SafeAreaView className={backgroundClassName}>
        <StatusBar style={isDark ? 'light' : 'dark'} />
        <View className="mx-auto w-full max-w-md flex-1 bg-gray-100 dark:bg-gray-900">
          <View className={headerClassName}>
            <View className="flex-row items-center justify-between px-4 py-6">
              <Pressable
                accessibilityLabel="Ir para o início"
                className="flex-row items-center gap-2"
                onPress={onHomePress}>
                <FontAwesome6 color="#0ea5e9" name="graduation-cap" size={42} />
                <Text className={`text-3xl font-bold ${textClassName}`}>StudyBuddy</Text>
              </Pressable>

              <View className="flex-row items-center gap-3">
                <Pressable
                  accessibilityLabel="Abrir notificações"
                  className={styles.headerIconButton(isDark)}
                  onPress={handleNotificationModal}>
                  <Ionicons color={iconColor} name="notifications" size={24} />
                  {mockNotifications.length > 0 && (
                    <View className="absolute right-0 top-0 h-5 min-w-5 items-center justify-center rounded-full bg-red-500 px-1">
                      <Text className="text-sm font-bold text-white">{mockNotifications.length}</Text>
                    </View>
                  )}
                </Pressable>

                <Pressable accessibilityLabel="Alternar tema" className={ isDark ? 'rounded-full bg-gray-700 px-2 py-4' : 'rounded-full bg-gray-100 px-2 py-4' } onPress={() => isDark ? setLightTheme() : setDarkTheme()} >
                  {isDark ? <MaterialIcons name="light-mode" size={24} color="white" /> : <MaterialIcons name="dark-mode" size={24} color="black" />}
                </Pressable>
              </View>
            </View>
          </View>

          {notificationModalVisible && (
            <SlideModalComponent visible={notificationModalVisible} onClose={() => setNotificationModalVisible(false)} title="Eventos de Hoje" >
              {mockNotifications.length === 0 ? (
                <Text className="text-black dark:text-white">
                  Nenhuma notificação hoje
                </Text>
              ) : (
                <View className="gap-3">
                  {mockNotifications.map((notification) => (
                    <NotificationCard key={notification.id} id={notification.id} title={notification.title} time={notification.time} description={notification.description}
                      onDelete={(id) => {
                        console.log('Excluir:', id);
                      }}
                    />
                  ))}
                </View>
              )}
            </SlideModalComponent>
          )}

          <ScrollView
            className="flex-1"
            contentContainerClassName="px-4 py-4 pb-28"
            showsVerticalScrollIndicator={false}>
            {children}
          </ScrollView>

          <View className={styles.bottomNav(isDark)} >
            <View className="mx-auto w-full max-w-md flex-row justify-around">
              {bottomTabs.map((tab) => {
                const isSelected = tab.id === selectedTab;
                const tabColor = isSelected ? '#0ea5e9' : isDark ? '#9ca3af' : '#6b7280';

                return (
                  <Pressable
                    accessibilityLabel={`Abrir aba ${tab.label}`}
                    accessibilityRole="tab"
                    accessibilityState={{ selected: isSelected }}
                    className="min-w-[64px] items-center px-2 py-3"
                    key={tab.id}
                    onPress={() => router.replace(`/${tab.id}`)}>
                    <FontAwesome6 color={tabColor} name={tab.icon} size={24} />
                    <Text className="mt-1 text-sm font-medium" style={{ color: tabColor }}>
                      {tab.label}
                    </Text>
                    <View className={isSelected ? styles.activeTabIndicator : styles.tabIndicator} />
                  </Pressable>
                );
              })}
            </View>
          </View>
        </View>
      </SafeAreaView>
    </SafeAreaProvider>
  );
};

const getTabFromRoute = (path: string, title: string): TabId => {
  const route = `${path} ${title}`.toLowerCase();

  if (route.includes('tecnica') || route.includes('tecnicas'))                return 'tecnicas';
  if (route.includes('ia')      || route.includes('inteligencia artificial')) return 'ia';
  if (route.includes('focus')   || route.includes('foco'))                    return 'focus';
  if (route.includes('config')  || route.includes('settings'))                return 'config';

  return 'home';
};

const styles = {
  activeTabIndicator: 'mt-2 h-[3px] w-8 rounded-full bg-sky-500',
  bottomNav: (isDark: boolean) =>
    isDark
      ? 'absolute inset-x-0 bottom-0 border-t border-gray-700 bg-gray-800 shadow-lg'
      : 'absolute inset-x-0 bottom-0 border-t border-gray-200 bg-white shadow-lg',
  headerIconButton: (isDark: boolean) =>
    isDark ? 'rounded-full bg-gray-700 px-2 py-4' : 'rounded-full bg-gray-100 px-2 py-4',
  tabIndicator: 'mt-2 h-[3px] w-8 rounded-full bg-transparent',
};
