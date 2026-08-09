import { useRouter } from 'expo-router';
import { useMemo, useState } from 'react';
import {
  ActivityIndicator,
  Alert,
  KeyboardAvoidingView,
  Linking,
  Platform,
  Pressable,
  ScrollView,
  Switch,
  Text,
  TextInput,
  View,
} from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';

const API_BASE_URL = 'http://localhost:8080';

type AuthUser = {
  id: number;
  email: string;
  name: string;
  createdAt: string;
};

type AuthResponse = {
  token: string;
  user: AuthUser;
};

export default function LoginScreen() {
  const router = useRouter();
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [rememberMe, setRememberMe] = useState(false);
  const [isDark, setIsDark] = useState(false);
  const [isSubmitting, setIsSubmitting] = useState(false);

  const canSubmit = useMemo(
    () => email.trim().length > 0 && password.length > 0 && !isSubmitting,
    [email, password, isSubmitting]
  );

  const handleLogin = async () => {
    if (!canSubmit) {
      Alert.alert('Campos obrigatórios', 'Informe email e senha para entrar.');
      return;
    }

    setIsSubmitting(true);

    try {
      const response = await fetch(`${API_BASE_URL}/auth/login`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          email: email.trim(),
          password,
          rememberMe,
        }),
      });

      const data = await response.json();

      if (!response.ok) {
        Alert.alert('Erro ao fazer login', data.error ?? 'Erro ao fazer login');
        return;
      }

      const { user } = data as AuthResponse;
      Alert.alert('Login realizado', `Bem-vindo(a), ${user.name || user.email}!`);
    } catch (error) {
      console.error('Erro ao conectar com a API:', error);
      Alert.alert('Erro de conexão', 'Não foi possível conectar com o servidor.');
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <SafeAreaView className={isDark ? 'flex-1 bg-gray-950' : 'flex-1 bg-gray-100'}>
      <KeyboardAvoidingView
        behavior={Platform.OS === 'ios' ? 'padding' : undefined}
        className="flex-1">
        <ScrollView
          contentContainerClassName="flex-grow justify-center px-4 py-8"
          keyboardShouldPersistTaps="handled">
          <View className="overflow-hidden rounded-3xl bg-sky-500 shadow-2xl">
            <View className={isDark ? 'bg-gray-800 p-8' : 'bg-white p-8'}>
              <View className="mb-2 flex-row justify-end">
                <Pressable
                  accessibilityLabel="Alternar tema"
                  className={
                    isDark ? 'rounded-full bg-gray-700 p-3' : 'rounded-full bg-gray-100 p-3'
                  }
                  onPress={() => setIsDark((current: boolean) => !current)}>
                  <Text className="text-lg">{isDark ? '◐' : '☾'}</Text>
                </Pressable>
              </View>

              <View className="mb-8 items-center">
                <Text
                  className={
                    isDark ? 'mb-4 text-5xl text-blue-400' : 'mb-4 text-5xl text-blue-600'
                  }>
                  🎓
                </Text>
                <Text
                  className={
                    isDark ? 'text-2xl font-bold text-white' : 'text-2xl font-bold text-gray-800'
                  }>
                  StudyBuddy
                </Text>
                <Text className={isDark ? 'mt-1 text-gray-300' : 'mt-1 text-gray-600'}>
                  Seu acompanhante de estudos
                </Text>
              </View>

              <View className="gap-4">
                <View>
                  <Text
                    className={
                      isDark
                        ? 'mb-1 text-sm font-medium text-gray-300'
                        : 'mb-1 text-sm font-medium text-gray-700'
                    }>
                    Email
                  </Text>
                  <TextInput
                    autoCapitalize="none"
                    autoComplete="email"
                    className={isDark ? inputStyles.dark : inputStyles.light}
                    editable={!isSubmitting}
                    keyboardType="email-address"
                    onChangeText={setEmail}
                    placeholder="seu@email.com"
                    placeholderTextColor={isDark ? '#9ca3af' : '#6b7280'}
                    value={email}
                  />
                </View>

                <View>
                  <Text
                    className={
                      isDark
                        ? 'mb-1 text-sm font-medium text-gray-300'
                        : 'mb-1 text-sm font-medium text-gray-700'
                    }>
                    Senha
                  </Text>
                  <TextInput
                    autoComplete="password"
                    className={isDark ? inputStyles.dark : inputStyles.light}
                    editable={!isSubmitting}
                    onChangeText={setPassword}
                    placeholder="••••••••"
                    placeholderTextColor={isDark ? '#9ca3af' : '#6b7280'}
                    secureTextEntry
                    value={password}
                  />
                </View>

                <View className="flex-row items-center justify-between">
                  <View className="flex-row items-center">
                    <Switch
                      onValueChange={setRememberMe}
                      thumbColor={rememberMe ? '#2563eb' : '#f3f4f6'}
                      trackColor={{ false: '#d1d5db', true: '#bfdbfe' }}
                      value={rememberMe}
                    />
                    <Text
                      className={
                        isDark ? 'ml-2 text-sm text-gray-300' : 'ml-2 text-sm text-gray-700'
                      }>
                      Lembrar de mim
                    </Text>
                  </View>

                  <Pressable
                    onPress={() =>
                      Alert.alert(
                        'Recuperar senha',
                        'A tela de recuperação será migrada em uma próxima etapa.'
                      )
                    }>
                    <Text className="text-sm font-medium text-blue-600">Esqueceu a senha?</Text>
                  </Pressable>
                </View>

                <Pressable
                  className={canSubmit ? buttonStyles.enabled : buttonStyles.disabled}
                  disabled={!canSubmit}
                  onPress={handleLogin}>
                  {isSubmitting ? (
                    <ActivityIndicator color="#ffffff" />
                  ) : (
                    <Text className="text-center font-medium text-white">Entrar</Text>
                  )}
                </Pressable>
              </View>

              <Text
                className={
                  isDark ? 'my-6 text-center text-gray-400' : 'my-6 text-center text-gray-500'
                }>
                ou
              </Text>

              <Pressable
                className={
                  isDark
                    ? 'w-full flex-row items-center justify-center gap-2 rounded-lg border border-gray-600 bg-gray-700 px-4 py-3'
                    : 'w-full flex-row items-center justify-center gap-2 rounded-lg border border-gray-300 bg-white px-4 py-3'
                }
                onPress={() => Linking.openURL(`${API_BASE_URL}/auth/google`)}>
                <Text className="text-lg">G</Text>
                <Text className={isDark ? 'text-white' : 'text-gray-800'}>Entrar com Google</Text>
              </Pressable>

              <View className="mt-6 flex-row justify-center">
                <Text className={isDark ? 'text-sm text-gray-400' : 'text-sm text-gray-600'}>
                  Não tem uma conta?{' '}
                </Text>
                <Pressable onPress={() => router.push('/register')}>
                  <Text className="text-sm font-medium text-blue-600">Cadastre-se</Text>
                </Pressable>
              </View>
            </View>
          </View>
        </ScrollView>
      </KeyboardAvoidingView>
    </SafeAreaView>
  );
}

const inputStyles = {
  light:
    'w-full rounded-lg border border-gray-300 bg-white px-4 py-3 text-gray-900 focus:border-blue-500',
  dark: 'w-full rounded-lg border border-gray-600 bg-gray-700 px-4 py-3 text-white focus:border-blue-500',
};

const buttonStyles = {
  enabled: 'w-full rounded-lg bg-blue-600 px-4 py-3',
  disabled: 'w-full rounded-lg bg-blue-300 px-4 py-3',
};
