import React from 'react';
import { Pressable, ScrollView, Switch, Text, TextInput, View } from 'react-native';

type AuthMode = 'login' | 'register' | 'forgot';
type Tab = 'dashboard' | 'notes' | 'calendar' | 'reminders' | 'materials';

const tabs: { id: Tab; label: string }[] = [
  { id: 'dashboard', label: 'Dashboard' },
  { id: 'notes', label: 'Notas' },
  { id: 'calendar', label: 'Calendário' },
  { id: 'reminders', label: 'Lembretes' },
  { id: 'materials', label: 'Materiais' },
];

const mockData = {
  notes: [
    { id: 1, title: 'Funções em JS', subject: 'Programação' },
    { id: 2, title: 'Segunda Guerra Mundial', subject: 'História' },
  ],
  events: [{ id: 1, title: 'Prova de Matemática', date: '12/08 08:00' }],
  reminders: [
    { id: 1, title: 'Revisar física', completed: false },
    { id: 2, title: 'Ler capítulo 3', completed: true },
  ],
  folders: ['Resumo', 'Aulas', 'Exercícios'],
};

export const ScreenContent: React.FC = () => {
  const [authMode, setAuthMode] = React.useState<AuthMode>('login');
  const [activeTab, setActiveTab] = React.useState<Tab>('dashboard');
  const [darkMode, setDarkMode] = React.useState(false);
  const [email, setEmail] = React.useState('');
  const [password, setPassword] = React.useState('');
  const [name, setName] = React.useState('');
  const [user, setUser] = React.useState<{ name: string; email: string } | null>(null);
  const [feedback, setFeedback] = React.useState('');

  const pendingReminders = mockData.reminders.filter((reminder) => !reminder.completed).length;

  const handleLogin = () => {
    if (!email || !password) {
      setFeedback('Preencha email e senha.');
      return;
    }

    setUser({ name: name || 'Estudante', email });
    setFeedback('');
  };

  const handleRegister = () => {
    if (!name || !email || !password) {
      setFeedback('Preencha nome, email e senha.');
      return;
    }

    setAuthMode('login');
    setFeedback('Conta criada com sucesso. Agora faça login.');
  };

  const handleForgotPassword = () => {
    if (!email) {
      setFeedback('Informe seu email.');
      return;
    }

    setAuthMode('login');
    setFeedback('Se o email existir, enviaremos instruções.');
  };

  if (!user) {
    return (
      <View className={`flex-1 justify-center p-6 ${darkMode ? 'bg-slate-900' : 'bg-slate-100'}`}>
        <View className="mb-4 flex-row items-center justify-end">
          <Text className={darkMode ? 'mr-2 text-slate-300' : 'mr-2 text-slate-700'}>
            Tema escuro
          </Text>
          <Switch value={darkMode} onValueChange={setDarkMode} />
        </View>

        <View className={`rounded-2xl p-6 ${darkMode ? 'bg-slate-800' : 'bg-white'}`}>
          <Text className={`mb-1 text-2xl font-bold ${darkMode ? 'text-white' : 'text-slate-900'}`}>
            StudyBuddy
          </Text>
          <Text className={`mb-5 ${darkMode ? 'text-slate-300' : 'text-slate-500'}`}>
            Seu acompanhante de estudos
          </Text>

          {authMode === 'register' && (
            <TextInput
              className={`mb-3 rounded-lg border px-3 py-2 ${darkMode ? 'border-slate-700 bg-slate-900 text-white' : 'border-slate-300 bg-white text-slate-900'}`}
              onChangeText={setName}
              placeholder="Nome"
              placeholderTextColor={darkMode ? '#94A3B8' : '#64748B'}
              value={name}
            />
          )}

          <TextInput
            autoCapitalize="none"
            className={`mb-3 rounded-lg border px-3 py-2 ${darkMode ? 'border-slate-700 bg-slate-900 text-white' : 'border-slate-300 bg-white text-slate-900'}`}
            keyboardType="email-address"
            onChangeText={setEmail}
            placeholder="Email"
            placeholderTextColor={darkMode ? '#94A3B8' : '#64748B'}
            value={email}
          />

          {authMode !== 'forgot' && (
            <TextInput
              className={`mb-4 rounded-lg border px-3 py-2 ${darkMode ? 'border-slate-700 bg-slate-900 text-white' : 'border-slate-300 bg-white text-slate-900'}`}
              onChangeText={setPassword}
              placeholder="Senha"
              placeholderTextColor={darkMode ? '#94A3B8' : '#64748B'}
              secureTextEntry
              value={password}
            />
          )}

          <Pressable
            className="mb-3 rounded-lg bg-sky-600 px-3 py-3"
            onPress={
              authMode === 'login'
                ? handleLogin
                : authMode === 'register'
                  ? handleRegister
                  : handleForgotPassword
            }>
            <Text className="text-center font-semibold text-white">
              {authMode === 'login'
                ? 'Entrar'
                : authMode === 'register'
                  ? 'Cadastrar'
                  : 'Enviar instruções'}
            </Text>
          </Pressable>

          {feedback ? (
            <Text
              className={`mb-3 text-sm ${feedback.includes('sucesso') ? 'text-emerald-600' : 'text-rose-500'}`}>
              {feedback}
            </Text>
          ) : null}

          <View className="flex-row justify-between">
            <Pressable onPress={() => setAuthMode('login')}>
              <Text className="text-sky-600">Login</Text>
            </Pressable>
            <Pressable onPress={() => setAuthMode('register')}>
              <Text className="text-sky-600">Cadastro</Text>
            </Pressable>
            <Pressable onPress={() => setAuthMode('forgot')}>
              <Text className="text-sky-600">Esqueci a senha</Text>
            </Pressable>
          </View>
        </View>
      </View>
    );
  }

  return (
    <View className={`flex-1 ${darkMode ? 'bg-slate-900' : 'bg-slate-100'}`}>
      <View className={`px-5 pb-4 pt-5 ${darkMode ? 'bg-slate-800' : 'bg-white'}`}>
        <View className="mb-2 flex-row items-center justify-between">
          <Text className={`text-xl font-bold ${darkMode ? 'text-white' : 'text-slate-900'}`}>
            StudyBuddy
          </Text>
          <View className="flex-row items-center gap-2">
            <Text className={darkMode ? 'text-slate-300' : 'text-slate-600'}>Escuro</Text>
            <Switch value={darkMode} onValueChange={setDarkMode} />
          </View>
        </View>
        <Text className={darkMode ? 'text-slate-300' : 'text-slate-600'}>Olá, {user.name}</Text>
      </View>

      <ScrollView className="flex-1 px-5 pt-4">
        <ScrollView horizontal showsHorizontalScrollIndicator={false}>
          <View className="mb-4 flex-row gap-2">
            {tabs.map((tab) => (
              <Pressable
                className={`rounded-full px-4 py-2 ${activeTab === tab.id ? 'bg-sky-600' : darkMode ? 'bg-slate-700' : 'bg-white'}`}
                key={tab.id}
                onPress={() => setActiveTab(tab.id)}>
                <Text
                  className={`font-medium ${activeTab === tab.id ? 'text-white' : darkMode ? 'text-slate-200' : 'text-slate-700'}`}>
                  {tab.label}
                </Text>
              </Pressable>
            ))}
          </View>
        </ScrollView>

        {activeTab === 'dashboard' && (
          <View className="gap-3">
            <View className={`rounded-xl p-4 ${darkMode ? 'bg-slate-800' : 'bg-white'}`}>
              <Text className={darkMode ? 'text-slate-300' : 'text-slate-500'}>
                Tempo de estudo
              </Text>
              <Text className={`text-2xl font-bold ${darkMode ? 'text-white' : 'text-slate-900'}`}>
                0h 0m
              </Text>
            </View>
            <View className={`rounded-xl p-4 ${darkMode ? 'bg-slate-800' : 'bg-white'}`}>
              <Text className={darkMode ? 'text-slate-300' : 'text-slate-500'}>
                Tarefas pendentes
              </Text>
              <Text className={`text-2xl font-bold ${darkMode ? 'text-white' : 'text-slate-900'}`}>
                {pendingReminders}/{mockData.reminders.length}
              </Text>
            </View>
          </View>
        )}

        {activeTab === 'notes' && (
          <View className="gap-3">
            {mockData.notes.map((note) => (
              <View
                className={`rounded-xl p-4 ${darkMode ? 'bg-slate-800' : 'bg-white'}`}
                key={note.id}>
                <Text className={`font-semibold ${darkMode ? 'text-white' : 'text-slate-900'}`}>
                  {note.title}
                </Text>
                <Text className={darkMode ? 'text-slate-300' : 'text-slate-500'}>
                  {note.subject}
                </Text>
              </View>
            ))}
          </View>
        )}

        {activeTab === 'calendar' && (
          <View className={`rounded-xl p-4 ${darkMode ? 'bg-slate-800' : 'bg-white'}`}>
            <Text className={`mb-2 font-semibold ${darkMode ? 'text-white' : 'text-slate-900'}`}>
              Próximos eventos
            </Text>
            {mockData.events.map((event) => (
              <Text className={darkMode ? 'text-slate-300' : 'text-slate-600'} key={event.id}>
                • {event.title} - {event.date}
              </Text>
            ))}
          </View>
        )}

        {activeTab === 'reminders' && (
          <View className="gap-3">
            {mockData.reminders.map((reminder) => (
              <View
                className={`rounded-xl p-4 ${darkMode ? 'bg-slate-800' : 'bg-white'}`}
                key={reminder.id}>
                <Text className={darkMode ? 'text-white' : 'text-slate-900'}>{reminder.title}</Text>
                <Text className={reminder.completed ? 'text-emerald-600' : 'text-amber-600'}>
                  {reminder.completed ? 'Concluído' : 'Pendente'}
                </Text>
              </View>
            ))}
          </View>
        )}

        {activeTab === 'materials' && (
          <View className="mb-5 gap-3">
            {mockData.folders.map((folder) => (
              <View
                className={`rounded-xl p-4 ${darkMode ? 'bg-slate-800' : 'bg-white'}`}
                key={folder}>
                <Text className={`font-medium ${darkMode ? 'text-white' : 'text-slate-900'}`}>
                  {folder}
                </Text>
              </View>
            ))}
          </View>
        )}

        <Pressable
          className="mb-8 mt-5 rounded-lg bg-rose-600 px-4 py-3"
          onPress={() => {
            setUser(null);
            setPassword('');
            setAuthMode('login');
          }}>
          <Text className="text-center font-semibold text-white">Sair</Text>
        </Pressable>
      </ScrollView>
    </View>
  );
};
