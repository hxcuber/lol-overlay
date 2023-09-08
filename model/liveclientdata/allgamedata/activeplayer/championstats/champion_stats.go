package championstats

type ChampionStats struct {
	AbilityHaste                 float64 `json:"abilityHaste"`
	AbilityPower                 float64 `json:"abilityPower"`
	Armor                        float64 `json:"armor"`
	ArmorPenetrationFlat         float64 `json:"armorPenetrationFlat"`
	ArmorPenetrationPercent      float64 `json:"armorPenetrationPercent"`
	AttackDamage                 float64 `json:"attackDamage"`
	AttackRange                  float64 `json:"attackRange"`
	AttackSpeed                  float64 `json:"attackSpeed"`
	BonusArmorPenetrationPercent float64 `json:"bonusArmorPenetrationPercent"`
	BonusMagicPenetrationPercent float64 `json:"bonusMagicPenetrationPercent"`
	CritChance                   float64 `json:"critChance"`
	CritDamage                   float64 `json:"critDamage"`
	CurrentHealth                float64 `json:"currentHealth"`
	HealShieldPower              float64 `json:"healShieldPower"`
	HealthRegenRate              float64 `json:"healthRegenRate"`
	LifeSteal                    float64 `json:"lifeSteal"`
	MagicLethality               float64 `json:"magicLethality"`
	MagicPenetrationFlat         float64 `json:"magicPenetrationFlat"`
	MagicPenetrationPercent      float64 `json:"magicPenetrationPercent"`
	MagicResist                  float64 `json:"magicResist"`
	MaxHealth                    float64 `json:"maxHealth"`
	MoveSpeed                    float64 `json:"moveSpeed"`
	Omnivamp                     float64 `json:"omnivamp"`
	PhysicalLethality            float64 `json:"physicalLethality"`
	PhysicalVamp                 float64 `json:"physicalVamp"`
	ResourceMax                  float64 `json:"resourceMax"`
	ResourceRegenRate            float64 `json:"resourceRegenRate"`
	ResourceType                 string  `json:"resourceType"`
	ResourceValue                float64 `json:"resourceValue"`
	SpellVamp                    float64 `json:"spellVamp"`
	Tenacity                     float64 `json:"tenacity"`
}
