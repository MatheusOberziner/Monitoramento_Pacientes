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
                <div class="col-12 col-md-5 q-pr-md">
                  <q-input
                    v-model="search"
                    outlined
                    :label="'Pesquisar por:'"
                    class=""
                    @keyup.enter="returnPatients()"
                  >
                    <template #append>
                      <q-btn-dropdown
                        flat
                        split
                        color="black"
                        icon="search"
                        @click="returnPatients()"
                      >
                        <q-list
                          dense
                          bordered
                          separator
                          style="width: 12em;"
                        >
                          <q-item
                            v-for="(filtro, index) in flitros"
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
                      @click="showUpdate(props.row)"
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

export default {
  name: 'IndexPage',
  setup () {
    const loading = ref(false)
    const filtroSelecionado = ref('nome')
    const search = ref('')

    return {
      filtroSelecionado,
      search,
      loading
    }
  },

  data () {
    return {
      filtros: [
        {
          label: 'Nome',
          value: 'nome'
        },
        {
          label: 'Idade',
          value: 'idade'
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
      rows: [
        { nome: 'Matheus', cpf: '999.999.999-99', sexo: 'Masculino', idade: '20', cidade: 'Rio do Sul' }
      ]
    }
  }
}
</script>
