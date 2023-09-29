<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated>
      <q-toolbar class="justify-between" style="background-color: #f7b7b7;">
          <div class="row q-px-md q-gutter-md">
            <div>
              <img src="../assets/logo.png" style="height: 80px; width: 80px;" />
            </div>
            <div class="self-center">
              <span class="text-h4 text-bold" style="color: black;">Monitoramento de Pacientes</span>
            </div>
          </div>

          <div class="row q-px-md q-gutter-md">
            <div>
              <q-btn round>
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  height="1.5em"
                  viewBox="0 0 24 24"
                  color="black"
                ><path
                  fill="currentColor"
                  d="M18 10a6 6 0 0 0-12 0v8h12v-8Zm2 8.667l.4.533a.5.5 0 0 1-.4.8H4a.5.5 0 0 1-.4-.8l.4-.533V10a8 8 0 1 1 16 0v8.667ZM9.5 21h5a2.5 2.5 0 0 1-5 0Z"
                /></svg>
                <q-badge
                  floating
                  color="red"
                  rounded
                />
              </q-btn>
            </div>

            <div>
              <q-btn round>
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  height="1.5em"
                  viewBox="0 0 24 24"
                  color="black"
                ><path
                  fill="currentColor"
                  d="M12 2C6.486 2 2 6.486 2 12v4.143C2 17.167 2.897 18 4 18h1a1 1 0 0 0 1-1v-5.143a1 1 0 0 0-1-1h-.908C4.648 6.987 7.978 4 12 4s7.352 2.987 7.908 6.857H19a1 1 0 0 0-1 1V18c0 1.103-.897 2-2 2h-2v-1h-4v3h6c2.206 0 4-1.794 4-4c1.103 0 2-.833 2-1.857V12c0-5.514-4.486-10-10-10z"
                /></svg>
              </q-btn>
            </div>

            <div>
              <q-btn-dropdown round icon="person" style="color: black;" />
            </div>
          </div>
      </q-toolbar>
    </q-header>

    <q-page-container>
      <q-page padding>
        <router-view />
      </q-page>
    </q-page-container>
  </q-layout>
</template>

<script>

export default {
  name: 'MainLayout',

  setup () {
  }
}
</script>
