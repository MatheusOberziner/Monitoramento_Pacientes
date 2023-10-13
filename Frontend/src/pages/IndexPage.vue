<template>
  <div class="row">
    <div class="col-12">
      <div class="col-12 col-md-12">
        <q-card>
          <q-card-section class="q-pb-none">
            <div class="row">
              <div class="col-12 col-md-4">
                <h4 class="q-px-md">
                  Lista de pacientes
                </h4>
              </div>

              <div class="col-12 col-md-8 q-pr-md row self-center justify-end">
                <div class="col-12 col-md-4 q-pr-md">
                  <q-input
                    v-model="search"
                    outlined
                    :label="'Filtrar por: ' + addLabel()"
                    class="uppercase"
                    @keyup.enter="returnPacientes()"
                  >
                    <template #append>
                      <q-btn-dropdown
                        flat
                        split
                        color="primary"
                        icon="search"
                        @click="returnPacientes()"
                      >
                        <q-list
                          dense
                          bordered
                          separator
                          style="width: 12em;"
                        >
                          <q-item
                            v-for="(filtro, index) in filtros"
                            :key="index"
                          >
                            <q-item-section side>
                              <q-radio
                                v-model="filtroSelecionado"
                                dense
                                checked-icon="task_alt"
                                unchecked-icon="panorama_fish_eye"
                                :label="filtro.label"
                                :val="filtro.value"
                              />
                            </q-item-section>
                          </q-item>
                        </q-list>
                      </q-btn-dropdown>
                    </template>
                  </q-input>
                </div>

                <div class="q-mt-sm">
                  <q-btn
                    label="Adicionar novo"
                    icon="add"
                    class="q-mb-md"
                    style="background-color: #f7b7b7;"
                    @click="showRegister()"
                  />
                </div>
              </div>
            </div>
          </q-card-section>
          <q-card-section>
            <div class="q-pa-md">
              <q-table
                v-model:pagination="serverPagination"
                :rows="rows"
                :columns="columns"
                :rows-per-page-options="[5,10,15]"
                :loading="loading"
              >
                <template #body-cell-actions="props">
                  <q-td :props="props" class="q-gutter-sm">
                    <q-btn
                      icon="info"
                      color="blue"
                      dense
                      size="sm"
                      style="width: 45px;"
                      @click="show(props.row)"
                    />
                  </q-td>
                </template>

                <template #no-data>
                  <q-icon
                    name="warning"
                    color="warning"
                    size="xs"
                    class="q-pa-sm"
                  />
                  <span style="font-size: 12px;">Nenhum dado encontrado!</span>
                </template>
              </q-table>
            </div>
          </q-card-section>
        </q-card>
      </div>
    </div>
  </div>
</template>
<script>
import { ref } from 'vue'
import { getPacientes } from 'boot/axios'
import RegisterPaciente from '../components/RegisterPaciente.vue'
export default {
  name: 'IndexPage',
  setup () {
    const loading = ref(false)
    const filtroSelecionado = ref('nome')
    const search = ref('')

    return {
      loading,
      filtroSelecionado,
      search
    }
  },
  data () {
    return {
      serverPagination: {
        page: 1,
        rowsPerPage: 10
      },
      filtros: [
        {
          label: 'Nome',
          value: 'nome'
        },
        {
          label: 'Cidade',
          value: 'cidade'
        }
      ],
      columns: [
        { name: 'nome', align: 'left', label: 'Nome', field: 'nome', sortable: true },
        { name: 'cpf', align: 'left', label: 'CPF', field: 'cpf', sortable: true },
        { name: 'sexo', align: 'left', label: 'Sexo', field: 'sexo', sortable: true },
        { name: 'idade', align: 'left', label: 'Idade', field: 'idade', sortable: true },
        { name: 'cidade', align: 'left', label: 'Cidade', field: 'cidade', sortable: true },
        { name: 'actions', align: 'left', label: 'Ações', field: 'actions' }
      ],
      rows: []
    }
  },
  created () {
    this.returnPacientes()
  },
  methods: {
    returnPacientes () {
      this.loading = true

      let nome = null, cidade = null
      switch (this.filtroSelecionado) {
        case 'nome':
          nome = this.search
          break
        case 'cidade':
          cidade = this.search
          break
        default:
          break
      }

      getPacientes({
        nome,
        cidade
      })
        .then(data => {
          this.rows = data
          this.rows.forEach(paciente => {
            paciente.idade = this.calculateAge(paciente.data_nascimento)
          })
          this.loading = false
        })
        .catch((error) => {
          this.loading = false
          this.rows = []
          this.$q.notify({
            message: 'Erro ao buscar pacientes...',
            caption: error,
            color: 'warning',
            position: 'top',
            actions: [
              { icon: 'close', color: 'white', round: true }
            ]
          })
        })
    },
    calculateAge (dateOfBorn) {
      const dob = new Date(dateOfBorn)
      const today = new Date()
      let age = today.getFullYear() - dob.getFullYear()
      const monthDiff = today.getMonth() - dob.getMonth()

      if (monthDiff < 0 || (monthDiff === 0 && today.getDate() < dob.getDate())) {
        age--
      }

      return age
    },
    addLabel () {
      const filtroSelecionado = this.filtroSelecionado

      const customFields = {
        cpf: 'CPF',
        cnpj: 'CNPJ',
        nome: 'Nome',
        cidade: 'Cidade'
      }

      return customFields[filtroSelecionado]
    },
    showRegister () {
      this.$q.dialog({
        component: RegisterPaciente,
        parent: this
      }).onOk(() => {
        this.returnPacientes()
      })
    }
  }
}
</script>
